function GetAllDoms(){
    $.getJSON( "/api/get/getalldomains", function( data ) {
        var items = [];
        $.each( data, function(key, val ) {
            items.push("<div class='row-sm-3 shadow-none p-3 mb-5 bg-light rounded'>"+key+") <b>" + val.Name +"</b><br>Uuid:"+val.Uuid+"<br>State:"+val.Status+"</div>" );
            console.log(key)
        });
        
        $( "<div/>", {
            "class": "col",
            html: items.join( "" )
        }).appendTo( ".container-fluid" );
        });
}

function HostMemInfo(){
    var value= $.ajax({
        dataType: "json", 
        url: '/api/get/hostmeminfo', 
        async: false
     });
     return value;
}
