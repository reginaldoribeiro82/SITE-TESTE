
Diagrama IHC (Protótipo) - Domínio Segurança

Este documento identifica os principais eventos de interação (ações do usuário) nas telas do domínio "Segurança" do aplicativo "Minha UFG" para os seguintes casos de uso:

Abrir chamado de segurança (UC_Seguranca_Abrir_1_1)

Consultar contatos de emergência (UC_Seguranca_Contatos_1_2)

Enviar localização para apoio (UC_Seguranca_Local_1_3)

Acionar Alerta (sino) (UC_Seguranca_Local_1_4)

Tela 1: Menu Principal "Minha UFG"

![Tela Principal](https://github.com/user-attachments/assets/59d531d7-3500-4563-b3fb-98cc4ecb5e51)


Esta é a tela principal do aplicativo, onde o usuário seleciona um dos 8 módulos. O fluxo se inicia aqui.

____________________________________________________________________________________________________________________________________________________________

Eventos de Interação Identificados:

Evento Nº 01: Acessar o Módulo Segurança

Evento: Ação ao Tocar

Componente: Ícone/Botão "Segurança" (no menu principal de 8 ícones).

Resultado Esperado: O aplicativo fecha a Tela Principal e abre a "Tela 2: Menu Segurança".
____________________________________________________________________________________________________________________________________________________________

Tela 2: Menu "Segurança"

![Modulo Segurança](https://github.com/user-attachments/assets/e6d4bef1-1be9-4db6-a9e7-8a16b598f1b4)


Esta é a tela principal do módulo, exibida após o Evento Nº 01. O usuário escolhe entre "Chamadas e Denúncias" ou "Emergência".
____________________________________________________________________________________________________________________________________________________________

Eventos de Interação Identificados:

Evento Nº 02: Clicar no botão "Chamadas e Denúncias"

Evento: Ação ao Tocar

Componente: Botão/Aba "CHAMADAS E DENÚNCIAS" (área de toque retangular).

Resultado Esperado: O conteúdo abaixo da aba é atualizado para mostrar a lista de categorias de denúncia (Agressão, Furto, etc.) abaixo da inscrição "PEÇA AJUDA".
____________________________________________________________________________________________________________________________________________________________

Evento Nº 03: Acionar o botão "Alerta de Segurança" (Sino)

Evento: Ação ao Tocar

Componente: Ícone de "Sino" (canto superior direito).

Resultado Esperado: O sistema operacional é acionado para, em paralelo, tocar um som de alerta e piscar o flash da câmera.
____________________________________________________________________________________________________________________________________________________________

Evento Nº 04: Clicar no botão "Emergência"

Evento: Ação ao Tocar 

Componente: Botão/Aba "EMERGÊNCIA".

Resultado Esperado: O conteúdo abaixo da aba é atualizado para mostrar a lista de contatos telefônicos de emergência (levando à "Tela 3").
____________________________________________________________________________________________________________________________________________________________

Evento Nº 05: Retornar ao Menu Principal (Tela 2)

Evento: Ação ao Tocar

Componente: Ícone "Voltar" (seta para a esquerda, no canto superior esquerdo).

Resultado Esperado: O aplicativo fecha a "Tela 2: Menu Segurança" e exibe novamente a "Tela 1: Menu Principal".
____________________________________________________________________________________________________________________________________________________________

Tela 3: Aba "Emergência" Ativa

![Tela Emergencia](https://github.com/user-attachments/assets/2ad693bb-4dc7-48c6-890f-75dea4169dec)

Esta tela é exibida após o usuário acionar o Evento Nº 04.
____________________________________________________________________________________________________________________________________________________________

Eventos de Interação Identificados:

Evento Nº 06: Clicar em um contato

Evento: Ação ao Tocar

Componente: Um item da lista (ex: "Cidade de Goiás 062 3521-2000").

Resultado Esperado: O aplicativo solicita ao Sistema Operacional a abertura do discador do celular, já preenchido com o número de telefone selecionado.
