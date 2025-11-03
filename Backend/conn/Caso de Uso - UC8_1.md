Diagrama IHC (Protótipo) - Domínio Apps UFG

Este documento identifica os principais eventos de interação (ações do utilizador) nas telas do domínio "Apps UFG" do aplicativo "Minha UFG" para o seguinte caso de uso:

Acessar Apps Externos (UC_Apps_8)

Tela 1: Menu Principal "Minha UFG"

![Menu Principal](https://github.com/user-attachments/assets/a5adb458-60b1-4908-91fd-1e98549f85b3)


Esta é a tela principal do aplicativo, onde o utilizador seleciona um dos 8 módulos. O fluxo inicia-se aqui.
_____________________________________________________________________________________________________________________________________________________________________

Eventos de Interação Identificados:

Evento Nº 01: Acessar ao Domínio Apps UFG

Evento: Ação ao clicar

Componente: Ícone/Botão "Apps UFG" (no menu principal de 8 ícones).

Resultado Esperado: O aplicativo fecha a Tela Principal e abre a "Tela 2: Apps UFG", carregando a lista de aplicações (ex: "RU UFG", "Inventário Patrimonial").
_____________________________________________________________________________________________________________________________________________________________________

Evento Nº 02: Retornar ao Menu Principal

Evento: Ação ao clicar

Componente: Ícone "Voltar" (seta para a esquerda, no canto superior esquerdo).

Resultado Esperado: O aplicativo fecha a "Tela 2: Apps UFG" e exibe novamente a "Tela 1: Menu Principal".
_____________________________________________________________________________________________________________________________________________________________________

Tela 2: Tela "Apps UFG"

![Imagem do WhatsApp de 2025-10-31 à(s) 07 36 31_1e8d7841](https://github.com/user-attachments/assets/02e3c26a-2f42-4db2-995f-ae403b5dfe09)


Esta tela é exibida após o Evento Nº 01. É aqui que o utilizador pode consultar a lista de apps relacionados e aceder às suas respectivas lojas para download.
_____________________________________________________________________________________________________________________________________________________________________

Eventos de Interação Identificados:

Evento Nº 03: Selecionar Link de App Externo

Evento: Ação ao clicar

Componente: O ícone de "link" ou "loja de aplicativo" (localizado à direita do nome de um app na lista).

Resultado Esperado: O aplicativo captura a URL associada a esse app e solicita ao usario se deseja baixar da loja de aplicativos, abrindo a tela 3
_____________________________________________________________________________________________________________________________________________________________________

Tela 3: Tela de Confirmação de Download

![Baixar app](https://github.com/user-attachments/assets/e03038f7-3847-442e-9b37-8073adbeaec6)

Esta tela é pop-up exibido sobre a "Tela 2" após o Evento Nº 02. Ela apresenta a mensagem: "O aplicativo não está instalado, deseja baixar da loja?" e os botões "Não" e "Sim".

_____________________________________________________________________________________________________________________________________________________________________


Eventos de Interação Identificados:

Evento Nº 04: Confirmar Download (Clicar "Sim")

Evento: onTap / onClick

Componente: Botão "Sim" no modal.

Resultado Esperado: O aplicativo fecha o modal, captura a URL associada ao app (selecionado no Evento 02) e solicita ao Sistema Operacional a abertura da loja de aplicativos (App Store/Play Store).

_____________________________________________________________________________________________________________________________________________________________________

Evento Nº 05: Cancelar Download (Clicar "Não")

Evento: onTap / onClick

Componente: Botão "Não" no modal.

Resultado Esperado: O aplicativo fecha o modal e o utilizador permanece na "Tela 2: Apps UFG".

