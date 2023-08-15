$('#formulario-cadastro').on('submit', criarUsuario);


function criarUsuario(evento) {
    evento.preventDefault();
    console.log("Dentro da função usuario!");   

    if ($('#senha').val() !== $('#confirmar-senha').val()) {
        alert("Ops... as senhas não coincidem");
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            senha: $('#senha').val()
        }
    }).done(function() {
        alert("Usuario cadastrado");
    }).fail(function(erro) {
        console.log(erro);
        alert("Usuario não cadastrado");
    });

}
   