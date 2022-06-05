package jgo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func ptr[T any](x T) *T {
	return &x
}

func TestUnmarshalFromString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    JSONObject
		wantErr bool
	}{
		{
			name: "valid JSON string",
			args: args{
				input: `{
					"name": "The Black Knight",
					"type": "Minion",
					"rarity": "Legendary",
					"cost": 6,
					"attack": 4.5,
					"taunts": [
						"Did I ever tell you what the definition of insanity is? Insanity is doing the exact... same fucking thing... over and over again expecting... shit to change... That. Is. Crazy.",
						"The first time somebody told me that, I dunno, I thought they were bullshitting me, so, I shot him. The thing is... He was right. And then I started seeing, everywhere I looked, everywhere I looked all these fucking pricks, everywhere I looked, doing the exact same fucking thing... over and over and over and over again thinking “This time is gonna be different' no, no, no please... This time is gonna be different.” Did I ever tell you the definition of…..Insanity?"
					],
					"skills": {
						"Attack": 1,
						"Block": 1
					}
				}`,
			},
			want: JSONObject{
				Values: map[string]JSONEntity{
					"name":   ptr(CreateJSONValue("The Black Knight")),
					"type":   ptr(CreateJSONValue("Minion")),
					"rarity": ptr(CreateJSONValue("Legendary")),
					"cost":   ptr(CreateJSONValue(6)),
					"attack": ptr(CreateJSONValue(4.5)),
					"taunts": &JSONArray{
						Values: []JSONEntity{
							ptr(CreateJSONValue("Did I ever tell you what the definition of insanity is? Insanity is doing the exact... same fucking thing... over and over again expecting... shit to change... That. Is. Crazy.")),
							ptr(CreateJSONValue("The first time somebody told me that, I dunno, I thought they were bullshitting me, so, I shot him. The thing is... He was right. And then I started seeing, everywhere I looked, everywhere I looked all these fucking pricks, everywhere I looked, doing the exact same fucking thing... over and over and over and over again thinking “This time is gonna be different' no, no, no please... This time is gonna be different.” Did I ever tell you the definition of…..Insanity?")),
						},
					},
					"skills": &JSONObject{
						Values: map[string]JSONEntity{
							"Attack": ptr(CreateJSONValue(1)),
							"Block":  ptr(CreateJSONValue(1)),
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalFromString(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				fmt.Println("\nGot: ")
				spew.Dump(got)
				fmt.Println("\n\nWant: ")
				spew.Dump(tt.want)
				t.Errorf("UnmarshalFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
