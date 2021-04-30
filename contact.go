package activecampaign

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

type Contact struct {
	Cdate               string `json:"cdate"`
	Email               string `json:"email"`
	Phone               string `json:"phone"`
	FirstName           string `json:"firstName"`
	LastName            string `json:"lastName"`
	Orgid               string `json:"orgid"`
	Orgname             string `json:"orgname"`
	Hash                string `json:"hash"`
	Deleted             string `json:"deleted"`
	Anonymized          string `json:"anonymized"`
	CreatedUtcTimestamp string `json:"created_utc_timestamp"`
	UpdatedUtcTimestamp string `json:"updated_utc_timestamp"`
	CreatedTimestamp    string `json:"created_timestamp"`
	UpdatedTimestamp    string `json:"updated_timestamp"`
	EmailEmpty          bool   `json:"email_empty"`
	ID                  string `json:"id"`
}

func UpsertContact(email string) (Contact, error) {
	log.Info("Verifying that contact with email " + email + " exists")

	var c Contact
	resp := ContactsResponse{}
	req, err := getRequest("api/3/contacts?email=" + email)
	if err != nil {
		return c, err
	}

	body, _ := ioutil.ReadAll(req.Body)

	err = json.Unmarshal(body, &resp)
	log.Info("Found " + resp.Meta.Total + " matching contacts")
	if err != nil {
		return c, err
	}

	if resp.Meta.Total == "0" {
		c, err := CreateContact(email)
		return c, err
	}

	for _, responseContact := range resp.Contacts {
		if responseContact.Email == email {
			c = responseContact
			break
		}
	}

	return c, nil
}

func CreateContact(email string) (Contact, error) {
	var c Contact
	type createPayload struct {
		Contact struct {
			Email string `json:"email"`
		} `json:"contact"`
	}
	payload := createPayload{}
	payload.Contact.Email = email

	payloadBytes, _ := json.Marshal(payload)
	req, err := postRequest("api/3/contacts", payloadBytes)
	if err != nil {
		return c, err
	}

	body, _ := ioutil.ReadAll(req.Body)
	type createResponse struct {
		Contact Contact `json:"contact"`
	}
	postResp := createResponse{}
	err = json.Unmarshal(body, &postResp)
	if err != nil {
		return c, err
	}

	log.Info("Created new contact with email " + postResp.Contact.Email)
	return postResp.Contact, nil
}

func (c *Contact) AddTag(tagID int) error {
	type addTagPayload struct {
		ContactTag struct {
			ContactID string `json:"contact"`
			TagID     string `json:"tag"`
		} `json:"contactTag"`
	}
	payload := addTagPayload{}
	payload.ContactTag.ContactID = c.ID
	payload.ContactTag.TagID = strconv.Itoa(tagID)

	payloadBytes, _ := json.Marshal(payload)
	req, err := postRequest("api/3/contactTags", payloadBytes)
	if err != nil {
		return err
	}

	if req.StatusCode != 201 {
		if req.StatusCode == 200 {
			log.Warn("Contact " + c.Email + " already has tag " + strconv.Itoa(tagID))
			return nil
		}
		return errors.New("Expected 201 response, got " + strconv.Itoa(req.StatusCode))
	}

	log.Info("Tagged " + c.Email + " with tag " + strconv.Itoa(tagID))

	return nil
}
