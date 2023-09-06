$('#nova-equipe').on("submit", criarEquipe);
$('#atualizar-equipe').on("click", atualizarEquipe);
$('.deletar-equipe').on("click", deletarEquipe);
$('.btn-danger').on('click', criarTarefaEquipe);
$('.concluir-tarefa-equipe').on("click", concluirTarefaDeEquipe);
$('.deletar-tarefa-equipe').on("click", deletarTarefaDeEquipe);
$('.editar-tarefa-equipe').on("click", editarTarefaDeEquipe);

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

function deletarEquipe(evento) {
    evento.preventDefault(); 

    Swal.fire({
        title: "Atenção!",
        text: "Deseja realmente excluir essa equipe? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;
    
   
    const elementoClicado = $(evento.target);
    const equipe = elementoClicado.closest('div');
    const equipeId = equipe.data('equipe-id');

    $.ajax({
        url: `/equipes/${equipeId}`,
        method: "DELETE"
    }).done(function() {
        equipe.fadeOut("slow", function() {
            $(this).remove();
        });    
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao excluir a equipe", "error");
    });
})
}

function criarTarefaEquipe(evento) {
    evento.preventDefault(); 

    // Obtenha o ID diretamente do botão de Publicar clicado
    const equipeId = $(this).data('equipe-id'); 

    $.ajax({
        url: `/equipes/${equipeId}/tarefas`,
        method: "POST",
        data: {
            tarefa: $('#tarefa-equipe').val(),
            observacao: $('#observacao-equipe').val(),
            prazo: $('#prazo-equipe').val(),
        }
    }).done(function() {
        window.location = `/equipes/${equipeId}/perfil`;
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao criar a tarefa!!!", "error");
    })
}

function concluirTarefaDeEquipe(evento) {
    evento.preventDefault(); 

    Swal.fire({
        title: "Atenção!",
        text: "Deseja realmente concluir essa tarefa?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;
   
        const elementoClicado = $(evento.target);
        const tarefa = elementoClicado.closest('div');
        const tarefaId = tarefa.data('tarefa-equipe-id');

        const equipeId = $('.col-xs-12.col-sm-12.col-md-7.col-lg-7.col-xl-7').data('equipe-id');

        elementoClicado.prop('disabled', true);

        $.ajax({
            url: `/equipes/${equipeId}/tarefas/${tarefaId}`,
            method: "DELETE"
        }).done(function() {
            tarefa.fadeOut("slow", function() {
                $(this).remove();
            });    
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao concluir a tarefa", "error");
        });
    });
}

function deletarTarefaDeEquipe(evento) {
    evento.preventDefault(); 

    Swal.fire({
        title: "Atenção!",
        text: "Deseja realmente excluir essa tarefa? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;
    
   
        const elementoClicado = $(evento.target);
        const tarefa = elementoClicado.closest('div');
        const tarefaId = tarefa.data('tarefa-equipe-id');

        elementoClicado.prop('disabled', true);

        const equipeId = $('.col-xs-12.col-sm-12.col-md-7.col-lg-7.col-xl-7').data('equipe-id');

        console.log("Tarefa ID:", tarefaId);

        $.ajax({
            url: `/equipes/${equipeId}/tarefas/${tarefaId}`,
            method: "DELETE"
        }).done(function() {
            tarefa.fadeOut("slow", function() {
                $(this).remove();
            });    
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao concluir a tarefa", "error");
        });
    });
}

function editarTarefaDeEquipe(evento) {
    
}

