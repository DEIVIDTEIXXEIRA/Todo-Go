$('#editar-usuario').on('submit', editar);

function editar(evento) {
evento.preventDefault();

$.ajax({
    url: "/editar-usuario",
    method: "PUT",
    data: {
        nome: $('#nome').val(),
        nick: $('#nick').val(),
        email: $('#email').val(),       
    }
}).done(function() {
    Swal.fire("Sucesso!", "Usuário atualizado com sucesso!", "success")
        .then(function() {
            window.location = "/perfil";
        });
}).fail(function() {
    Swal.fire("Ops...", "Erro ao atualizar o usuário!", "error");
});
}