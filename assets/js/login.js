$('#formulario-login').on('submit', fazerLogin);

function fazerLogin(evento){
    evento.preventDefault();


    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            senha: $('#senha').val()
        }
    }).done(function(){ // 200 201 204
        window.location = "/home";
    }).fail(function(){ // 400 500 
        Swal.fire("Ops...", "User or password incorrect!", "error");
    });

}