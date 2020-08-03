package main
import(
	"database/sql"
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"strconv"
	"io/ioutil"
	"strings"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)
func handleRequests(){
	myApp:=mux.NewRouter().StrictSlash(true)
myApp.HandleFunc("/add", BasicAuthMiddleware(http.HandlerFunc(Addemp))).Methods("POST")
myApp.HandleFunc("/update",BasicAuthMiddleware(Updateemp)).Methods("POST")
myApp.HandleFunc("/search",Searchemp).Methods("POST")
myApp.HandleFunc("/list",Listemp).Methods("POST")
myApp.HandleFunc("/delete",BasicAuthMiddleware(Deleteemp)).Methods("POST")
myApp.HandleFunc("/restore",BasicAuthMiddleware(Restoreemp)).Methods("POST")



log.Fatal(http.ListenAndServe(":8080",myApp))
}
func main(){
	handleRequests()
}
///////////////////////////////////////
type Emp struct{
	Name string `json:"name"`
	Department string `json:"department"`
	Address string `json:"address"`
	Skills []string `json:"skills"`
}
func Addemp(w http.ResponseWriter,r *http.Request){
	fmt.Println("addemp")
	var e Emp
	log.Println("BODY IS",r.Body)
	
	err:=json.NewDecoder(r.Body).Decode(&e)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		//w.WriteHeader(400)
		log.Fatal(err)
		return
	}
	log.Println("dept is :",e.Department)
	log.Println("name is :",e.Name)
	if e.Name==""{
		log.Printf("name is compulsory:")
		w.WriteHeader(401)
	w.Write([]byte("name is compulsory:"))
		return
	}
	
	
	name:=e.Name
	dept:=e.Department
	addr:=e.Address
	var skills string
	 for _,k:=range e.Skills{
		 skills+=k+","
	 }
	 skills=skills[:len(skills)-1]
	db:=dbConn()
	addempbody,err:=db.Prepare("INSERT INTO employee(name,dept,addr,skills,active) VALUES(?,?,?,?,?)")

	if err != nil {
	panic(err.Error())
	}
	addempbody.Exec(name, dept,addr,skills,true)
	log.Println("INSERT: Name: " + name + " | dept: " + dept+ " | addr: " + addr+ " | skills: " + skills)
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	 
	json.NewEncoder(w).Encode(e)
	defer db.Close()

}
func dbConn()(db *sql.DB){
	dbDriver := "mysql"
    dbUser := "root"
    dbPass := "sudhadargainya"
    dbName := "empdb"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}
////////////////////////////////////////////////////
type Deleteempstr struct{
	Empid int `json:"empid"`
	PermanentlyDelete bool `json:"permanentlyDelete"`
	}
func Deleteemp(w http.ResponseWriter,r *http.Request){
	db:=dbConn()
	fmt.Println("deleteemp")
	var deleteemp Deleteempstr 
	
	err:=json.NewDecoder(r.Body).Decode(&deleteemp)
   	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) 
		return
	}
	
	var id int
	id=deleteemp.Empid
	fmt.Println("id is ....",id)
	log.Println("BODY PARAMS:=",strconv.FormatBool(deleteemp.PermanentlyDelete)+"id: "+strconv.Itoa(id))
	if deleteemp.PermanentlyDelete==true {
		log.Println("employee has to be permanently deleted")
	delForm, err := db.Prepare("DELETE FROM employee WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(id)
    log.Println("DELETE")
	}else{
	log.Println("employee has to be deactivated")
		delForm, err := db.Prepare("UPDATE employee SET active=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(false,id)
	 
	}
defer db.Close()
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(200)
json.NewEncoder(w).Encode(id)	
}
//////////////////////////////////////////////
type Listempstr struct{
	Empid int `json:"empid"`
	Name string `json:"name"`
	Department string `json:"department"`
	Address string `json:"address"`
	Skills []string `json:"skills"`
}
func Listemp(w http.ResponseWriter,r *http.Request){
	
	db:=dbConn()
fmt.Println("Listemp")
var listemp Listempstr 
reqBody, err := ioutil.ReadAll(r.Body)
   if err != nil {
	log.Printf("Body read error, %v", err)
	w.WriteHeader(500) 
	return
}
err=json.Unmarshal(reqBody, &listemp)
var id int
if listemp.Empid!=0{
	id=listemp.Empid
}

var name string
if listemp.Name!=""{
name=listemp.Name
}
var dept string
if listemp.Department!=""{
dept=listemp.Department
}
var addr string
if listemp.Address!=""{
addr=listemp.Address
}


var skills string
if len(listemp.Skills)!=0{
for _,k:=range listemp.Skills{
//skills=append(skills,updateemp.Skills)
skills+=k+","
}
}
log.Println("BODY PARAMS:=",dept+ " "+addr+" "+skills+" "+"name: "+name+"id :"+strconv.Itoa(id))


if  name=="" && dept=="" && addr=="" && skills=="" && id==0{
	log.Println("********listing all employee***********")
	selDB, err :=db.Query("SELECT * FROM employee")
	if err != nil {
		panic(err.Error())
		
	}
	emp := Listempstr{}
	res := []Listempstr{}
	var ssid int
	var  ssdept,ssaddr,ssskilss,ssname string
	var sactive bool
for selDB.Next() {
	
	err = selDB.Scan(&ssid,&ssname, &ssdept, &ssaddr,&ssskilss,&sactive)
	
	if err != nil {
		panic(err.Error())
	}
	emp.Empid = ssid
	emp.Name = ssname
	emp.Department = ssdept
	emp.Address=ssaddr
	emp.Skills=strings.Split(ssskilss,",")
	if(sactive==true){
	res = append(res, emp)
	}
	fmt.Println("res...............",res)
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(res)	

}else {
log.Println("********listing particular employees based upon particular fields********",id,name,dept)
selDB, err := db.Query("SELECT * FROM employee WHERE id=? OR name=? OR dept=? ",id,name,dept)
if err != nil {
	panic(err.Error())
	
}
emp := Listempstr{}
res := []Listempstr{}
var ssid int
	var  ssname,ssdept,ssaddr,ssskilss string
	var sactive bool
for selDB.Next() {
	
	err = selDB.Scan(&ssid,&ssname, &ssdept, &ssaddr,&ssskilss,&sactive)
	
	if err != nil {
		panic(err.Error())
	}
	emp.Empid = ssid
	emp.Name = name
	emp.Department = ssdept
	emp.Address=ssaddr
	emp.Skills=strings.Split(ssskilss,",")
	if(sactive==true){
	res = append(res, emp)
	}
}

w.Header().Set("Content-Type", "application/json")
log.Println("********* RESULT IS ***********",res)
json.NewEncoder(w).Encode(res)	

}



defer db.Close()
}
//////////////////////////////////////////////////////
type Restoreempstr struct{
	Empid int `json:"empid"`
	
	}
func Restoreemp(w http.ResponseWriter,r *http.Request){
	db:=dbConn()
	fmt.Println("deleteemp")
	var restoremep Restoreempstr 
	
	err:=json.NewDecoder(r.Body).Decode(&restoremep)
   	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) 
		return
	}
	
	var id int
	id=restoremep.Empid
	fmt.Println("id is ....",id)
	log.Println("BODY PARAMS:=id: "+strconv.Itoa(id))
	
	log.Println("employee has to be Activated")
		delForm, err := db.Prepare("UPDATE employee SET active=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(true,id)
	
	
defer db.Close()

}
/////////////////////////////////////////
type Searchempstr struct{
	Empid int `json:"empid"`
	Name string `json:"name"`
	Department string `json:"department"`
	Address string `json:"address"`
	Skills []string `json:"skills"`
}

func Searchemp(w http.ResponseWriter,r *http.Request){

	db:=dbConn()
fmt.Println("searchemp")
var searchemp Searchempstr 
reqBody, err := ioutil.ReadAll(r.Body)
   if err != nil {
	log.Printf("Body read error, %v", err)
	w.WriteHeader(500) 
	return
}
err=json.Unmarshal(reqBody, &searchemp)
var name string
if searchemp.Name!=""{
name=searchemp.Name
}
var dept string
if searchemp.Department!=""{
dept=searchemp.Department
}
var addr string
if searchemp.Address!=""{
addr=searchemp.Address
}


var skills string
if len(searchemp.Skills)!=0{
for _,k:=range searchemp.Skills{

skills+=k+","
}
skills=skills[:len(skills)-1] //////justvadd
}
log.Println("BODY PARAMS:=",dept+ " "+addr+" "+skills+" "+"name: "+name)
selDB, err := db.Query("SELECT id,name,dept,addr,skills,active FROM employee WHERE name=? OR dept=? OR addr=? OR skills=?",name,dept,addr,skills)
if err != nil {
	panic(err.Error())
}

emp := Searchempstr{}
res := []Searchempstr{}
for selDB.Next() {
	var ssid int
	var  ssdept,ssaddr,ssskilss,ssname string
	var ssactive bool
	err = selDB.Scan(&ssid,&ssname, &ssdept, &ssaddr,&ssskilss,&ssactive)
		if err != nil {
		panic(err.Error())
	}
	emp.Empid = ssid
	emp.Name = ssname
	emp.Department = ssdept
	emp.Address=ssaddr
	emp.Skills=strings.Split(ssskilss,",")
	if ssactive==true{
	res = append(res, emp)
	}
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(res)	

defer db.Close()
}
/////////////////////////////////////////////////////
type UpdateEmpstr struct{
	Empid int `json:"empid"`
	Department string `json:"department"`
	Address string `json:"address"`
	Skills []string `json:"skills"`
}
func Updateemp(w http.ResponseWriter,r *http.Request){
	db:=dbConn()
	fmt.Println("updateemp")
	var updateemp UpdateEmpstr 
	reqBody, err := ioutil.ReadAll(r.Body)
   	if err != nil {
		log.Printf("Body read error, %v", err)
		w.WriteHeader(500) 
		return
	}
	err=json.Unmarshal(reqBody, &updateemp)
	var dept string
	if updateemp.Department!=""{
	dept=updateemp.Department
	}
	var addr string
	if updateemp.Address!=""{
	addr=updateemp.Address
	}
	var id int
	
	id=updateemp.Empid
	fmt.Println("id is ....",id)
	
	var skills string
	if len(updateemp.Skills)!=0{
	for _,k:=range updateemp.Skills{

skills+=k+","
	}
	skills=skills[:len(skills)-1]
	}
	log.Println("BODY PARAMS:=",dept+ " "+addr+" "+skills+" "+"id: "+strconv.Itoa(id))
	updtemp, err := db.Prepare("UPDATE employee SET dept=?, addr=?, skills=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	updtemp.Exec(dept, addr, skills,id)
	log.Println("UPDATE: dept: " + dept + " | addr: " + addr+ " | skills: " + skills ) 

defer db.Close()

}
////////////////////////////////////////////
//basic auth
func BasicAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		fmt.Println("username: ", user)
		fmt.Println("password: ", pass)
		if !ok || !checkUsernameAndPassword(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		handler(w, r)
	}
}

func checkUsernameAndPassword(username, password string) bool {
	return username == "rupa" && password == "jain"
}