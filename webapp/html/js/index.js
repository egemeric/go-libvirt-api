function GetAllDoms(){
    var xhttp = new XMLHttpRequest();
    xhttp.open("GET","http://127.0.0.1:9000/api/get/getalldomains",false);
    xhttp.send(null);
    return xhttp.responseText;
}