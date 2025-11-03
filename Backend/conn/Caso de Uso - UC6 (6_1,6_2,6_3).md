Diagrama IHC (Protótipo) - Domínio Rádio UFG

Este documento identifica os principais eventos de interação (ações do utilizador) nas telas do domínio "Rádio" do aplicativo "Minha UFG" para os seguintes casos de uso:

Consultar programação semanal (UC_Radio_Prog_6_1)

Ouvir rádio ao vivo (UC_Radio_Ouvir_6_2)

Ver informações de programas (UC_Radio_Info_6_3)



Tela 1: Menu Principal "Minha UFG"

![Tela Principal](https://github.com/user-attachments/assets/68c64d2b-9168-434d-b1d9-7520505e0d1e)


Esta é a tela principal do aplicativo, onde o utilizador seleciona um dos 8 módulos. O fluxo inicia-se aqui.
_________________________________________________________________________________________________________________________________________________________________

Eventos de Interação Identificados:

Evento Nº 01: Acessar ao Domínio Rádio

Evento: Ação ao Tocar

Componente: Ícone/Botão "Rádio" (no menu principal de 8 ícones).

Resultado Esperado: O aplicativo fecha a Tela Principal e abre a "Tela 2: Rádio UFG", carregando a programação do dia atual.
_________________________________________________________________________________________________________________________________________________________________

Tela 2: Tela "Rádio UFG"

![Tela Radio UFG](https://github.com/user-attachments/assets/f62b8107-3abb-4197-aa42-cdd0fbd8e3e2)


Esta tela é exibida após o Evento Nº 01. É aqui que o utilizador pode consultar a programação e ouvir a rádio ao vivo.

_________________________________________________________________________________________________________________________________________________________________

Eventos de Interação Identificados:

Evento Nº 02: Selecionar Dia da Semana (Consultar Programação)

Evento: Ação ao Tocar

Componente: Botão/Aba de um dia da semana (ex: "DOM", "SEG", "TER").

Resultado Esperado: A lista de programas abaixo é atualizada para mostrar a grade de programação do dia selecionado, incluindo horários, nomes dos programas e descrições.
_________________________________________________________________________________________________________________________________________________________________

Evento Nº 03: Ouvir Rádio Ao Vivo

Evento: Ação ao Tocar

Componente: Botão "Play" (ícone de triângulo, localizado na parte inferior da tela).

Resultado Esperado: O aplicativo inicia o streaming de áudio da rádio ao vivo. O ícone do botão muda para "Pause".
_________________________________________________________________________________________________________________________________________________________________

Evento Nº 04: Retornar ao Menu Principal

Evento: Ação ao Tocar

Componente: Ícone "Voltar" (seta para a esquerda, no canto superior esquerdo).

Resultado Esperado: O aplicativo fecha a "Tela 2: Rádio UFG" e exibe novamente a "Tela 1: Menu Principal".
_________________________________________________________________________________________________________________________________________________________________
