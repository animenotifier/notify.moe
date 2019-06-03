package arn

import "errors"

// PushSubscriptions is a list of push subscriptions made by a user.
type PushSubscriptions struct {
	UserID UserID              `json:"userId"`
	Items  []*PushSubscription `json:"items"`
}

// Add adds a subscription to the list if it hasn't been added yet.
func (list *PushSubscriptions) Add(subscription *PushSubscription) error {
	if list.Contains(subscription.ID()) {
		return errors.New("PushSubscription " + subscription.ID() + " has already been added")
	}

	subscription.Created = DateTimeUTC()

	list.Items = append(list.Items, subscription)

	return nil
}

// Remove removes the subscription ID from the list.
func (list *PushSubscriptions) Remove(subscriptionID string) bool {
	for index, item := range list.Items {
		if item.ID() == subscriptionID {
			list.Items = append(list.Items[:index], list.Items[index+1:]...)
			return true
		}
	}

	return false
}

// Contains checks if the list contains the subscription ID already.
func (list *PushSubscriptions) Contains(subscriptionID string) bool {
	for _, item := range list.Items {
		if item.ID() == subscriptionID {
			return true
		}
	}

	return false
}

// Find returns the subscription with the specified ID, if available.
func (list *PushSubscriptions) Find(id string) *PushSubscription {
	for _, item := range list.Items {
		if item.ID() == id {
			return item
		}
	}

	return nil
}

// GetPushSubscriptions ...
func GetPushSubscriptions(id string) (*PushSubscriptions, error) {
	obj, err := DB.Get("PushSubscriptions", id)

	if err != nil {
		return nil, err
	}

	return obj.(*PushSubscriptions), nil
}
