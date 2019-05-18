package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func PutUser(w http.ResponseWriter, r *http.Request){
    user := UserDet{}
    pUser := UserDet{}
    status := Confirm{
        Status: "Error",
        Text: "",
    }
    err := json.NewDecoder(r.Body).Decode(&user);
    if(err!=nil){
        panic(err)
    }
    fmt.Println("Got these :",user.Username,user.Password)
    db := GetDB();

    //check if already exists in users
    _,errr := db.Exec("SELECT * FROM users WHERE reg = $1",user.Username)
    if(errr != nil){
        //he no exist
        //check if he exists in pwi
        errp := db.QueryRow("SELECT reg,pwd,name,hostel FROM pwi WHERE reg = $1",user.Username).Scan(&pUser.Username,&pUser.Password,&pUser.Name,&pUser.Hostel)
        if(errp != nil){
            //he no exists in pwi
            status.Status = "Wrong"
            goto EXIT
        }else{
            //he exist in pwi
            //insert him into ours
            _,erra := db.Exec("INSERT INTO users VALUES ($1,$2,$3,$4,'false');",user.Username,user.Password,pUser.Name,pUser.Hostel)
            if(erra!=nil){
                //insert him to users failed
                status.Status = "Fault"
                goto EXIT
            }else{
                //inserted him into users
                status.Status = "OK"
                goto EXIT
            }
        }
    }else{
        //he exist
        status.Status = "Exists"
        goto EXIT
    }
    EXIT:
    statusJson,errj := json.Marshal(status)
    if errj!=nil{
        panic(errj)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
