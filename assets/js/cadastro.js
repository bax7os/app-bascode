$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento){
    evento.preventDefault();


    if($('#senha').val() !== $('#confirmar-senha').val()){
        Swal.fire("Ops...", "The passwords do not match!", "error");
        return;
    }
    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function() {
        Swal.fire("Sucess!", "User created with success!", "success")
            .then(function() {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),
                        senha: $('#senha').val()
                    }
                }).done(function() {
                    window.location = "/home";
                }).fail(function() {
                    Swal.fire("Ops...", "Failed to authenticate the user!", "error");
                })
            })
    }).fail(function() {
        Swal.fire("Ops...", "Failed to create the user!", "error");
    });

}