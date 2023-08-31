$('#nova-equipe').on("submit", criarEquipe);
$('#atualizar-equipe').on("click", atualizarEquipe);

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

function atualizarEquipe(evento) {
    $(this).prop('disabled', true);

    const equipeId = $(this).data('equipe-id');

    $.ajax({
        url: `/equipes/${equipeId}`,
        method: "PUT",
        data: {
            nome: $('#nome').val(),
            descricao: $('#descricao').val()
        }
    }).done(function() {
        Swal.fire( 
            'Sucesso',
            'Equipe atualizada com sucesso!',
            'success')
            .then(function() {
                window.location = "/equipe";
            })
        }).fail(function() {
            Swal.fire("Ops...", "Falha em editar a equipe!!", "error");
    }).always(function() {
        $('#atualizar-equipe').prop('disabled', false)
    });
    
}