package lucca

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/bureaux-a-partager/lucca-leaves-slack-reminder/lib/utils"
)

type LeavesResponse struct {
	Data struct {
		Item []struct {
			IsAM        bool `json:"isAM"`
			LeavePeriod struct {
				Owner struct {
					Name string `json:"name"`
				} `json:"owner"`
				EndsOn string `json:"endsOn"`
				EndsAM bool   `json:"endsAM"`
			} `json:"leavePeriod"`
		} `json:"items"`
	} `json:"data"`
}

type Employee struct {
	Name      string
	Morning   bool
	Afternoon bool
	AllDay    bool
	EndDate   string
}

// List employees (en build good objects)
func (lr *LeavesResponse) ListEmployees() []Employee {
	list := []Employee{}
	var e Employee
	var index int

	for _, item := range lr.Data.Item {
		exists := false
		for itemIndex, itemList := range list {
			if itemList.Name == item.LeavePeriod.Owner.Name {
				e = itemList
				exists = true
				index = itemIndex
			}
		}

		if !exists {
			dt := strings.Split(item.LeavePeriod.EndsOn, "T")
			e = Employee{
				Name:      item.LeavePeriod.Owner.Name,
				Morning:   false,
				Afternoon: false,
				AllDay:    false,
				EndDate:   dt[0],
			}
			list = append(list, e)
			index = len(list) - 1
		}

		if item.IsAM {
			list[index].Morning = true
		}

		if !item.IsAM {
			list[index].Afternoon = true
		}

		if list[index].Morning && list[index].Afternoon {
			list[index].AllDay = true
		}
	}

	return list
}

// Format employees list
func FormatEmployees(employees []Employee) string {
	now := time.Now().Format("2006-01-02")

	output := ""

	for _, employee := range employees {
		output = output + "- " + employee.Name + " est absent(e)"
		if employee.EndDate == now {
			output = output + " aujourd'hui"
			if employee.AllDay {
				output = output + " toute la journée"
			} else if employee.Morning {
				output = output + " le matin"
			}
		} else {
			output = output + " jusqu'au " + utils.DateIsoToFr(employee.EndDate)
			if employee.AllDay {
				output = output + " inclus"
			} else if employee.Morning {
				output = output + " à midi"
			}
		}
		output = output + "\n"
	}

	return output
}

func GetLeaves(token string) LeavesResponse {
	now := time.Now().Format("2006-01-02")
	url := "https://bap.ilucca.net/api/v3/leaves?leavePeriod.ownerId=greaterthan,0&date=" + now + "&fields=isAM,leavePeriod[owner.name,endsOn,endsAM]"

	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "lucca application="+token)

	res, _ := client.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	lr := LeavesResponse{}

	err := json.Unmarshal([]byte(string(body)), &lr)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return lr
}
