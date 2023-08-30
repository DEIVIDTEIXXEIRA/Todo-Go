CREATE DATABASE IF NOT EXISTS todo;
USE todo;

DROP TABLE IF EXISTS usuarios_equipe;
DROP TABLE IF EXISTS tarefas_equipe;
DROP TABLE IF EXISTS equipes;
DROP TABLE IF EXISTS tarefas;
DROP TABLE IF EXISTS usuarios;


CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL
) ENGINE=InnoDB;


CREATE TABLE tarefas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tarefa VARCHAR(100) NOT NULL,
    observacao VARCHAR(100) NOT NULL,
    prazo VARCHAR(50) NOT NULL,
    autor_id INT NOT NULL,
    FOREIGN KEY (autor_id) REFERENCES usuarios(id) ON DELETE CASCADE, 
    autor_nick VARCHAR(50) NOT NULL,
    FOREIGN KEY (autor_nick) REFERENCES usuarios(nick) ON DELETE CASCADE
) ENGINE=InnoDB;


CREATE TABLE equipes (
    id INT AUTO_INCREMENT PRIMARY KEY, 
    nome VARCHAR(50) NOT NULL UNIQUE,
    descricao VARCHAR(100) NOT NULL,

    autor_id INT NOT NULL,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE tarefas_equipe (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tarefa VARCHAR(100) NOT NULL,
    observacao VARCHAR(100) NOT NULL,
    prazo VARCHAR(50) NOT NULL,
    autor_id INT NOT NULL,
    equipes_id INT NOT NULL,
    FOREIGN KEY (autor_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    FOREIGN KEY (equipes_id) REFERENCES equipes(id) ON DELETE CASCADE
) ENGINE=InnoDB;




CREATE TABLE usuarios_equipe (
    PRIMARY KEY (equipes_id, usuario_id),
    
    equipes_id INT NOT NULL,
    FOREIGN KEY (equipes_id)
    REFERENCES equipes(id)
    ON DELETE CASCADE,
    
    usuario_id INT NOT NULL,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    usuario_nick VARCHAR(50) NOT NULL,
    FOREIGN KEY (usuario_nick)
    REFERENCES usuarios(nick)
    ON DELETE CASCADE
) ENGINE=InnoDB;
