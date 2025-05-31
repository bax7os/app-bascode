$("#nova-publicacao").on("submit", criarPublicacao);
$(document).on("click", ".curtir-publicacao", curtirPublicacao);
$(document).on("click", ".descurtir-publicacao", descurtirPublicacao);
$(".deletar-publicacao").on("click", deletarPublicacao);
$("#atualizar-publicaco").on("click", editarPublicacao);
function criarPublicacao(evento) {
  evento.preventDefault();
  $.ajax({
    url: "/publicacoes",
    method: "POST",
    data: {
      titulo: $("#titulo").val(),
      conteudo: $("#conteudo").val(),
    },
  })
    .done(function () {
      // 200 201 204
      window.location = "/home";
    })
    .fail(function () {
        Swal.fire("Ops...", "Error creating the post!", "error");
    });
}

function curtirPublicacao(evento) {
  evento.preventDefault();

  const elementoClicado = $(evento.target);
  const publicacaoId = elementoClicado.closest("div").data("publicacao-id");
  elementoClicado.prop("disabled", true);
  $.ajax({
    url: `/publicacoes/${publicacaoId}/curtir`,
    method: "POST",
  })
    .done(function () {
      // 200 201 204
      const contadorDeCurtidas = elementoClicado.next("span");
      const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
      contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

      elementoClicado.addClass("descurtir-publicacao");
      elementoClicado.addClass("text-danger");
      elementoClicado.removeClass("curtir-publicacao");
    })
    .fail(function () {
      Swal.fire("Ops...", "Failed to like the post!", "error");
    })
    .always(function () {
      elementoClicado.prop("disabled", false);
    });
}

function deletarPublicacao(evento) {
  evento.preventDefault();

  Swal.fire({
    title: "Are you sure?",
    text: "You won't be able to revert this!",
    icon: "warning",
    showCancelButton: true,
    confirmButtonColor: "#3085d6",
    cancelButtonColor: "#d33",
    confirmButtonText: "Yes, delete it!",
  }).then((result) => {
    if (!result.value) {
      return;
    }
    const elementoClicado = $(evento.target);
  
    elementoClicado.prop("disabled", true);
    const publicacao = elementoClicado.closest("div");
    const publicacaoId = publicacao.data("publicacao-id");
    $.ajax({
      url: `/publicacoes/${publicacaoId}`,
      method: "DELETE",
    })
      .done(function () {
       publicacao.fadeOut("slow", function () {
         $(this).remove();
       })
        
      })
      .fail(function () {
     
        Swal.fire("Ops...", "Failed to delete the post!", "error");
      })
      .always(function () {
        elementoClicado.prop("disabled", false);
      });
  });


 
}

function editarPublicacao() {
  $(this).prop("disabled", true);

  const publicacaoId = $(this).data("publicacao-id");

  $.ajax({
    url: `/publicacoes/${publicacaoId}`,
  method: "PUT",
  data: {
    "titulo": $("#titulo").val(),
    "conteudo": $("#conteudo").val()
  }
  }).done(function() {
   
   Swal.fire("Sucess!", "Post edited with success!", "success").then(function() {
     window.location = "/home";
   })
  }).fail(function() {
    Swal.fire("Ops...", "Failed to update the post!", "error");
  }).always(function() {
    $("#atualizar-publicaco").prop("disabled", false);
  });


}

function descurtirPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest("div").data("publicacao-id");
    elementoClicado.prop("disabled", true);
    $.ajax({
      url: `/publicacoes/${publicacaoId}/descurtir`,
      method: "POST",
    })
      .done(function () {
        // 200 201 204
        const contadorDeCurtidas = elementoClicado.next("span");
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);
  
        elementoClicado.removeClass("descurtir-publicacao");
        elementoClicado.removeClass("text-danger");
        elementoClicado.addClass("curtir-publicacao");
      })
      .fail(function () {
        // 400 500
    
        Swal.fire("Ops...", "Failed to like the post!", "error");
      })
      .always(function () {
        elementoClicado.prop("disabled", false);
      });
}