$("#parar-seguir").on("click", pararDeSeguir);
$("#seguir").on("click", seguir);
$(".editar-usuario").on("submit", editar);
$("#editar-senha").on("submit", atualizarSenha);
$("#deletar-usuario").on("click", deletarUsuario);
function pararDeSeguir(){
    const usuarioID = $(this).data("usuario-id")
    console.log(usuarioID);
    $(this).prop('disabled', true);
    $.ajax({
        url: `/usuarios/${usuarioID}/parar-de-seguir`,
        method: "POST"
    }).done(function(){
        window.location= `/usuarios/${usuarioID}`;
    }).fail(function(){
        Swal.fire("Ops...", "Failed to unfollow the user!", "error");
        $("parar-de-seguir").prop('disabled', false);
    });
        
}

function seguir(){
    const usuarioID = $(this).data("usuario-id")
    console.log(usuarioID);
    $(this).prop('disabled', true);
    $.ajax({
        url: `/usuarios/${usuarioID}/seguir`,
        method: "POST"
    }).done(function(){
        window.location= `/usuarios/${usuarioID}`;
    }).fail(function(){
        Swal.fire("Ops...", "Failed to follow the user!", "error");
        $("#seguir").prop('disabled', false);
    });
}

function editar(evento){
    evento.preventDefault();
    $.ajax({
        url: "/editar-usuario",
        method: "PUT",
        data: {
            nome: $("#nome").val(),
            email: $("#email").val(),
            nick: $("#nick").val()
            
        }
    }).done(function(){
        Swal.fire("Sucess!", "User edited with success!", "success").then(function(){
            window.location = "/perfil";
        });
    }).fail(function(){
        Swal.fire("Ops...", "Failed to edit the user!", "error");
    });
}

function atualizarSenha(evento){
    evento.preventDefault();

    if($("#nova-senha").val() !== $("#confirmar-senha").val()){
        Swal.fire("Ops...", "The passwords do not match!", "error");
        return;
    }

    $.ajax({
        url: "/atualizar-senha",
        method: "POST",
        data: {
            antiga: $("#senha-atual").val(),
            nova: $("#nova-senha").val()
        }
    }).done(function(){
        Swal.fire("Sucess!", "Password edited with success!", "success").then(function(){
            window.location = "/perfil";
        });
    }).fail(function(){
        Swal.fire("Ops...", "Failed to edit the password!", "error");
    });

}

function deletarUsuario(){
    Swal.fire({
        title: 'Are you sure?',
        text: "You won't be able to revert this!",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Yes, delete my account!'
    }).then((result) => {
        if (result.isConfirmed) {
            $.ajax({
                url: "/deletar-usuario",
                method: "DELETE"
            }).done(function(){
                Swal.fire("Sucess!", "User deleted with success!", "success").then(function(){
                    window.location = "/logout";
                });
            }).fail(function(){
                Swal.fire("Ops...", "Failed to delete the user!", "error");
            });
        }
    })
}