package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	// route path folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// routing
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/form-project", formAddProject).Methods("GET")
	route.HandleFunc("/form-editproject/{index}", formEditProject).Methods("GET")
	route.HandleFunc("/detail-project/{index}", detailProject).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("POST")
	route.HandleFunc("/edit-project", editProject).Methods("POST")
	route.HandleFunc("/delete-project/{index}", deleteProject).Methods("GET")

	fmt.Println("server running on port 5050")
	// menjalankan server
	http.ListenAndServe("localhost:5050", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Description-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return
	}
	// menampilka data dari dataProject yang di input data add project
	response := map[string]interface{}{
		"Projects": dataProject,
	}

	// w.Write([]byte("add project"))
	// w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, response)
}

func formAddProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Description-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/addproject.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

// membuat teplate (dto = data transformation object)
type Project struct {
	ProjectName string
	Description string
	StartDate   string
	EndDate     string
}

// mengambl data yang di push dari newproject dalam bentuk array
var dataProject = []Project{
	{
		ProjectName: "Dummy ProjectName",
		Description: "Dummy Description",
	},
	{
		ProjectName: "Dummy ProjectName 1",
		Description: "Dummy Description 1",
	},
}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("ProjectName: " + r.PostForm.Get("ProjectName"))
	// fmt.Println("Description: " + r.PostForm.Get("description"))

	title := r.PostForm.Get("projectName")
	content := r.PostForm.Get("description")
	startDate := r.PostForm.Get("startDate")
	endDate := r.PostForm.Get("endDate")

	// menyimpan data dari form ke object
	newProject := Project{
		ProjectName: title,
		Description: content,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	// push data dari onject ke dataproject
	dataProject = append(dataProject, newProject)
	fmt.Println(dataProject)
	// redirect ke halaman index setelah button di klik
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func detailProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Description-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/detailproject.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return
	}
	// menangkap index params dari url
	index, _ := strconv.Atoi(mux.Vars(r)["index"])
	// index, _ := mux.Vars(r)["index"]

	var projectDetail = Project{}

	// menampung data object dan menyesuaikan i dari data lopp dan index data param
	for i, data := range dataProject {
		if i == index {
			projectDetail = Project{
				ProjectName: data.ProjectName,
				Description: data.Description,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
			}
		}
	}
	// menampilkan data
	data := map[string]interface{}{
		"Project": projectDetail,
	}

	tmpl.Execute(w, data)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Description-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	index, _ := strconv.Atoi(mux.Vars(r)["index"])
	// fmt.Println(index)
	dataProject = append(dataProject[:index], dataProject[index+1:]...)

	http.Redirect(w, r, "/", http.StatusFound)
}

func formEditProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Description-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/editproject.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return
	}
	index, _ := strconv.Atoi(mux.Vars(r)["index"])
	fmt.Println(index)

	var projectEdit = Project{}

	// menampung data object dan menyesuaikan i dari data lopp dan index data param
	for i, data := range dataProject {
		if i == index {
			projectEdit = Project{
				ProjectName: data.ProjectName,
				Description: data.Description,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
			}
		}
	}
	// menampilkan data
	data := map[string]interface{}{
		"Project": projectEdit,
	}
	fmt.Println(data)

	tmpl.Execute(w, data)
}
func editProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	title := r.PostForm.Get("projectName")
	content := r.PostForm.Get("description")
	startDate := r.PostForm.Get("startDate")
	endDate := r.PostForm.Get("endDate")

	// menyimpan data dari form ke object
	newProject := Project{
		ProjectName: title,
		Description: content,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	// push data dari onject ke dataproject
	dataProject = append(dataProject, newProject)
	fmt.Println(dataProject)
	// redirect ke halaman index setelah button di klik
	http.Redirect(w, r, "/", http.StatusFound)

}
