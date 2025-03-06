# Desafio Itaú
Desafio de _live coding_ para o cargo de Engenheiro Sênior no Itaú. Consistiu no desenvolvimento de uma API REST que
implementasse o famoso jogo Jokenpô (do japonês じゃんけんぽん, romanizado como _jankenpon_), também conhecido no Brasil
como Pedra, Papel e Tesoura.
# Como deve funcionar o jogo?
O jogador deve escolher uma opção de jogada.
O computador faz uma escolha aleatória.
A API deve decidir o vencedor com base nas regras do jogo.
# Premissas
O jogador deve ser capaz de inserir somente as opções "pedra", "papel" ou "tesoura".
Qualquer outra entrada deve ser considerada inválida.
O computador deve escolher somente entre "pedra", "papel" ou "tesoura".
A API deve exibir as escolhas do jogador e do computador.
A API deve informar o vencedor da rodada ou se houve empate.
# Regras do jogo
Pedra vence Tesoura
Tesoura vence Papel
Papel vence Pedra
Caso as escolhas sejam iguais, o resultado é um empate.
# Rodando o jogo
No Terminal, vá até o repositório e execute `go build main.go`.
Com a aplicação rodando, envie a requisição abaixo pelo Terminal:
`curl -X POST http://localhost:8080/play \
     -H "Content-Type: application/json" \
     -d '{"move": "paper"}' | jq
`
O parâmetro `move` pode receber os valores `rock`, `paper` e `scissors`. A aplicação deverá devolver um JSON contendo o
movimento do jogador que o usuário acabou de inserir, o movimento aleatoriamente escolhido pelo computador e o resultado da partida.
Desenvolveu-se também TUs para os métodos `getWinner` e `handlerPlay`. Os testes podem ser executados através do comando
`go test -v`. O código foi escrito em inglês por habitualidade do desenvolvedor.
