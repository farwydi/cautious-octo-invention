package main

import (
	"fmt"
	"github.com/looplab/fsm"
	"log"
)

func main() {
	var iq = 100
	machine := fsm.NewFSM(
		"",
		fsm.Events{
			{
				Name: "go_result",
				Src: []string{
					"question_2_answer_1",
					"question_2_answer_2",
					"question_2_answer_3",
				},
				Dst: "result",
			},
			{Name: "go_hello", Src: []string{""}, Dst: "hello"},
			{Name: "go_save_name", Src: []string{"hello"}, Dst: "save_name"},
			{
				Name: "go_question_1",
				Src: []string{
					"save_name",
				},
				Dst: "question_1",
			},
			{Name: "1", Src: []string{"question_1"}, Dst: "question_1_answer_1"},
			{Name: "2", Src: []string{"question_1"}, Dst: "question_1_answer_2"},
			{Name: "3", Src: []string{"question_1"}, Dst: "question_1_answer_3"},
			{Name: "1", Src: []string{"question_2"}, Dst: "question_2_answer_1"},
			{Name: "2", Src: []string{"question_2"}, Dst: "question_2_answer_2"},
			{Name: "3", Src: []string{"question_2"}, Dst: "question_2_answer_3"},
			{
				Name: "go_question_2",
				Src: []string{
					"question_1_answer_1",
					"question_1_answer_2",
					"question_1_answer_3",
				},
				Dst: "question_2",
			},
		},
		fsm.Callbacks{
			"hello": func(event *fsm.Event) {
				_, err := fmt.Println(localization[event.Dst])
				if err != nil {
					log.Fatalln(err)
				}
			},
			"enter_question_2": func(event *fsm.Event) {
				_, err := fmt.Println(localization["question_2"])
				if err != nil {
					log.Fatalln(err)
				}
				_, err = fmt.Println("1:\t", localization["question_2_answer_1"])
				if err != nil {
					log.Fatalln(err)
				}
				_, err = fmt.Println("2:\t", localization["question_2_answer_2"])
				if err != nil {
					log.Fatalln(err)
				}
				_, err = fmt.Println("3:\t", localization["question_2_answer_3"])
				if err != nil {
					log.Fatalln(err)
				}
			},
			"enter_question_2_answer_1": func(event *fsm.Event) {
				iq -= 10
			},
			"enter_question_2_answer_2": func(event *fsm.Event) {
				iq += 10
			},
			"enter_question_2_answer_3": func(event *fsm.Event) {
				iq -= 10
			},
			"enter_question_1": func(event *fsm.Event) {
				_, err := fmt.Println(localization["question_1"])
				if err != nil {
					log.Fatalln(err)
				}
				_, err = fmt.Println("1:\t", localization["question_1_answer_1"])
				if err != nil {
					log.Fatalln(err)
				}
				_, err = fmt.Println("2:\t", localization["question_1_answer_2"])
				if err != nil {
					log.Fatalln(err)
				}
				_, err = fmt.Println("3:\t", localization["question_1_answer_3"])
				if err != nil {
					log.Fatalln(err)
				}
			},
			"enter_question_1_answer_1": func(event *fsm.Event) {
				iq += 10
			},
			"enter_question_1_answer_2": func(event *fsm.Event) {
				iq -= 10
			},
			"enter_question_1_answer_3": func(event *fsm.Event) {
				iq -= 10
			},
			"enter_result": func(event *fsm.Event) {
				_, err := fmt.Printf("You IQ: %d\n", iq)
				if err != nil {
					log.Fatalln(err)
				}
			},
			"before_go_save_name": func(event *fsm.Event) {
				localization["save_name"] = fmt.Sprintf(localization["save_name"], event.Args...)
			},
		},
	)

	err := machine.Event("go_hello")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		state := machine.Current()
		switch state {
		case "result":
			return
		case "question_2_answer_1", "question_2_answer_2", "question_2_answer_3":
			err := machine.Event("go_result")
			if err != nil {
				log.Fatalln(err)
			}
		case "question_1_answer_1", "question_1_answer_2", "question_1_answer_3":
			err := machine.Event("go_question_2")
			if err != nil {
				log.Fatalln(err)
			}
		case "save_name":
			err := machine.Event("go_question_1")
			if err != nil {
				log.Fatalln(err)
			}
		default:
			read := ""
			_, err := fmt.Print("You answer: ")
			if err != nil {
				log.Fatalln(err)
			}
			_, err = fmt.Scanln(&read)
			if err != nil {
				log.Fatalln(err)
			}
			switch state {
			case "hello":
				err = machine.Event("go_save_name", read)
				if err != nil {
					log.Fatalln(err)
				}
			default:
				err = machine.Event(read)
				if err != nil {
					_, err := fmt.Println(localization["err"])
					if err != nil {
						log.Fatalln(err)
					}
				}
			}
		}
	}

	//switch machine.Current() {
	//case "hello":
	//	_, err :=fmt.Println(localization[machine.Current()])
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	event := ""
	//	_, err = fmt.Scanln(&event)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	err = machine.Event(event)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//case "question_1":
	//case "question_2":
	//
	//}
	//
	//fmt.Println(machine.Current())
	//
	//err := machine.Event("open")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(machine.Current())
	//
	//err = machine.Event("close")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(machine.Current())
}
