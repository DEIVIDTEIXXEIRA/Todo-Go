$('#nova-tarefa').on("submit", criarTarefa);
$('.concluir-tarefa').on("click", concluirTarefa);

function criarTarefa(evento) {
    evento.preventDefault(); 

    $.ajax({
        url: "/tarefas",
        method: "POST",
        data: {
            tarefa: $('#tarefa').val(),
            observacao: $('#observacao').val(),
            prazo: $('#prazo').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        alert("erro ao criar tarefa!!");
    })
}

function concluirTarefa(evento) {
    evento.preventDefault(); 
   
    const elementoClicado = $(evento.target);
    const tarefa = elementoClicado.closest('div');
    const tarefaId = tarefa.data('tarefa-id');

    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/tarefas/${tarefaId}`,
        method: "DELETE"
    });
    
    tarefa.fadeOut("slow", function() {
        $(this).remove();
    });
    
} 