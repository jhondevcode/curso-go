package main

import "fmt"

/* ***** Relacion de uno a uno ***** */
type User struct {
	name  string
	email string
	state bool
}

type Student struct {
	user User
	code string
}

/* ***** Relacion de uno a muchos ***** */
type Video struct {
	title  string
	course Course
}

type Course struct {
	title  string
	videos []Video
}

func main() {
	/* Relacion de 1 a 1 */
	var alex User = User{
		name:  "Alex",
		email: "alex@gmail.com",
		state: true,
	}

	var roel User = User{
		name:  "Roel",
		email: "roel@gmail.com",
		state: true,
	}

	var student1 Student = Student{
		user: alex,
		code: "0X46567816",
	}

	fmt.Println(roel)
	fmt.Println(student1)

	/* Relacion de 1 a varios */

	var video01 Video = Video{
		title: "01-Intoduccion",
	}

	var video02 Video = Video{
		title: "02-Instalation",
	}

	var course Course = Course{
		title:  "Curso de Go",
		videos: []Video{video01, video02},
	}

	video01.course = course
	video02.course = course

	for _, video := range course.videos {
		fmt.Println(video.title)
	}
}
