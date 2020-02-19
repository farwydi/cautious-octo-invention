package main

import (
	"fmt"
	"github.com/looplab/fsm"
	"log"
)

func main() {
	machine := fsm.NewFSM(
		"hello",
		fsm.Events{
			{Name: "go_hello", Src: []string{}, Dst: "hello"},
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
			{Name: "4", Src: []string{"question_1"}, Dst: "question_1_answer_4"},
			{Name: "1", Src: []string{"question_2"}, Dst: "question_2_answer_1"},
			{Name: "2", Src: []string{"question_2"}, Dst: "question_2_answer_2"},
			{Name: "3", Src: []string{"question_2"}, Dst: "question_2_answer_3"},
			{Name: "4", Src: []string{"question_2"}, Dst: "question_2_answer_4"},
			{
				Name: "go_question_2",
				Src: []string{
					"question_2_answer_1",
					"question_2_answer_2",
					"question_2_answer_3",
					"question_2_answer_4",
				},
				Dst: "question_2",
			},
		},
		fsm.Callbacks{
			"after_question_1": func(event *fsm.Event) {
				_, err := fmt.Println("1:\t", localization["question_1_answer_1"])
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
			"before_go_save_name": func(event *fsm.Event) {
				localization["save_name"] = fmt.Sprintf(localization["save_name"], event.Args...)
			},
		},
	)

	for {
		state := machine.Current()
		_, err := fmt.Println(localization[state])
		if err != nil {
			log.Fatalln(err)
		}
		switch state {
		case "save_name":
			err = machine.Event("go_question_1")
			if err != nil {
				log.Fatalln(err)
			}
			continue
		case "question_1":
			_, err := fmt.Println("1:\t", localization["question_1_answer_1"])
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
		}
		read := ""
		_, err = fmt.Print("You answer: ")
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
