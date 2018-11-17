package main 

import (
    "database/sql"
    "log"
    "net/http"
    "text/template"
    _ "github.com/go-sql-driver/mysql"
)

// student object
type Student struct {
    Id    int
    Name  string
    Course string
}


// db connection
func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "godb"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

// frontend directory
var tmpl = template.Must(template.ParseGlob("frontend/*"))

// function to query all Student, use in HandleFunc
func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM Student ORDER BY id DESC")
    if err != nil {
        panic(err.Error())
    }
    student := Student{}
    res := []Student{}
    for selDB.Next() {
        var id int
        var name, course string
        err = selDB.Scan(&id, &name, &course)
        if err != nil {
            panic(err.Error())
        }
        student.Id = id
        student.Name = name
        student.Course = course
        res = append(res, student)
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

// function to display Student detail
func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM Student WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    student := Student{}
    for selDB.Next() {
        var id int
        var name, course string
        err = selDB.Scan(&id, &name, &course)
        if err != nil {
            panic(err.Error())
        }
        student.Id = id
        student.Name = name
        student.Course = course
    }
    tmpl.ExecuteTemplate(w, "Show", student)
    defer db.Close()
}

// display template for new student
func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

// function edit student detail
func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM Student WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    student := Student{}
    for selDB.Next() {
        var id int
        var name, course string
        err = selDB.Scan(&id, &name, &course)
        if err != nil {
            panic(err.Error())
        }
        student.Id = id
        student.Name = name
        student.Course = course
    }
    tmpl.ExecuteTemplate(w, "Edit", student)
    defer db.Close()
}


// function to insert new Student
func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        course := r.FormValue("course")
        insForm, err := db.Prepare("INSERT INTO Student(name, course) VALUES(?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, course)
        log.Println("INSERT: Name: " + name + " | Course: " + course)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

// update record
func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        course := r.FormValue("course")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE Student SET name=?, course=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, course, id)
        log.Println("UPDATE: Name: " + name + " | Course: " + course)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

// delete record
func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    student := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM Student WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(student)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}


// entry point
func main() {
    log.Println("Server started on: http://localhost:8080")
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe(":8080", nil)
}