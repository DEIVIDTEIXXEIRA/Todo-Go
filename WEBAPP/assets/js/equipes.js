$('#nova-equipe').on("submit", criarEquipe);

function criarEquipe(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/equipe",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            observacao: $('#descricao').val(),
        }
    }).done(function() {
        window.location = "/equipe";
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao criar a equipe!!!", "error");
    })
}