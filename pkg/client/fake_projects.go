package client

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
)
import projectapi "github.com/openshift/origin/pkg/project/api"

type FakeProjects struct {
	Fake *Fake
}

func (c *FakeProjects) List(label labels.Selector, field fields.Selector) (*projectapi.ProjectList, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "list-projects"}, &projectapi.ProjectList{})
	return obj.(*projectapi.ProjectList), err
}

func (c *FakeProjects) Get(name string) (*projectapi.Project, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "get-project"}, &projectapi.Project{})
	return obj.(*projectapi.Project), err
}

func (c *FakeProjects) Create(project *projectapi.Project) (*projectapi.Project, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "create-project", Value: project}, &projectapi.Project{})
	return obj.(*projectapi.Project), err
}

func (c *FakeProjects) Update(project *projectapi.Project) (*projectapi.Project, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "update-project"}, &projectapi.Project{})
	return obj.(*projectapi.Project), err
}

func (c *FakeProjects) Delete(name string) error {
	c.Fake.Actions = append(c.Fake.Actions, FakeAction{Action: "delete-project", Value: name})
	return nil
}
