# Cheat Sheet para Colaboradores de Kubernetes

Una lista de recursos comunes al contribuir a Kubernetes, consejos, trucos y
mejores prácticas comúnmente utilizadas dentro del proyecto Kubernetes. Es un
resumen o referencia rápida de información útil para hacer tu experiencia de contribución en GitHub
mejor.

**Tabla de Contenidos**
- [Recursos Útiles](#recursos-útiles)
  - [Empezar](#empezar)
  - [SIGs y Otros Grupos](#sigs-y-otros-grupos)
  - [Comunidad](#comunidad)
  - [Flujo de Trabajo](#flujo-de-trabajo)
  - [Pruebas](#pruebas)
  - [Alias de Correo Importantes](#alias-de-correo-importantes)
  - [Otros Enlaces Útiles](#otros-enlaces-útiles)
- [Comunicación Efectiva en GitHub](#comunicación-efectiva-en-github)
  - [Cómo Ser Excelente Entre Sí](#cómo-ser-excelente-entre-sí)
    - [Ejemplos de Buena/Mala Comunicación](#ejemplos-de-buenamala-comunicación)
- [Presentar una Contribución](#presentar-una-contribución)
  - [Firmar el CLA](#firmar-el-cla)
  - [Abrir y Responder Issues](#abrir-y-responder-issues)
    - [Crear un Issue](#crear-un-issue)
    - [Responder a un Issue](#responder-a-un-issue)
  - [Abrir un Pull Request](#abrir-un-pull-request)
    - [Crear un Pull Request](#crear-un-pull-request)
    - [Ejemplo de Descripción de un Pull Request](#ejemplo-de-descripción-de-un-pull-request)
    - [Solucionar Problemas de un Pull Request](#solucionar-problemas-de-un-pull-request)
  - [Labels](#labels)
- [Trabajar Localmente](#trabajar-localmente)
  - [Estrategia de Branch](#estrategia-de-branch)
    - [Agregar Upstream](#agregar-upstream)
    - [Mantener tu Fork Sincronizado](#mantener-tu-fork-sincronizado)
    - [Combinar Commits (Squashing Commits)](#squashing-commits)
   
---
## Recursos Útiles

### Empezar

- [Curso para Colaboradores] - **NUEVO** - El curso de E-Learning para colaboradores de Kubernetes!
- [Guía para Colaboradores] - Guía sobre cómo comenzar a contribuir al proyecto
  Kubernetes.
- [Guía del Desarrollador] - Guía para contribuir código directamente al proyecto
  Kubernetes.
- [Guía de Seguridad y Divulgación] - Guía para informar vulnerabilidades
  y el proceso de liberación de seguridad.

### SIGs y Otros Grupos

- [Lista Maestra de Grupos][SIGs]

### Comunidad

- [Calendario] - Visualiza todos los eventos de la Comunidad de Kubernetes (reuniones de SIG/WG,
  eventos, etc.).
- [kubernetes-dev] - La lista de correo de desarrollo de Kubernetes.
- [Foro de Kubernetes] - Foro oficial de Kubernetes.
- [Canales de Slack] - Slack oficial de Kubernetes.
- [Stack Overflow] - Un lugar para hacer preguntas como usuario final de Kubernetes.
- [Canal de YouTube] - Canal oficial de la comunidad de Kubernetes.

### Flujo de Trabajo

- [Prow] - Sistema de CI/CD de Kubernetes.
- [Tide] - Plugin de Prow que administra fusiones y pruebas. [Panel de Tide]
- [Comandos de Bots] - Comandos utilizados para interactuar con los Bots de Kubernetes (ejemplos:
  `/cc`, `/lgtm` y `/retest`).
- [GitHub Labels] - Lista de labels utilizados en todo el proyecto Kubernetes.
- [Búsqueda de Código de Kubernetes], mantenida por [@dims].

### Pruebas

- [Prow] - Sistema de CI/CD de Kubernetes.
- [Test Grid] - Visualiza pruebas históricas y su información asociada.
- [Panel de Triage] - Agrega problemas similares para una mejor
  resolución de problemas.

### Alias de Correo Importantes

- community@kubernetes.io - Envía un correo a alguien del equipo de la comunidad
  (SIG Contributor Experience) acerca de un problema de comunidad.
- conduct@kubernetes.io - Contacta al comité de Código de Conducta,
  lista de correo privada.
- github@kubernetes.io - Envía un correo al [Equipo de Administración de GitHub] de forma privada,
  para asuntos sensibles.
- steering@kubernetes.io - Envía un correo al comité directivo. Dirección pública
  con archivo público.
- steering-private@kubernetes.io - Envía un correo al comité directivo de forma privada,
  para asuntos sensibles.
- social@cncf.io - Contacta al equipo social de la CNCF; blog, cuenta de Twitter
  y otras propiedades sociales.

### Otros Enlaces Útiles

- [Estadísticas de Desarrolladores] - Visualiza estadísticas de desarrolladores para todos los
  proyectos gestionados por la CNCF.
- [Publicación de Parches de Kubernetes] - Programación e información de contacto del equipo
  para lanzamientos de parches de Kubernetes.

---

## Comunicación Efectiva en GitHub

### Cómo ser Excelente entre Sí

Como primer paso, familiarízate con el [Código de Conducta].

#### Ejemplos de Buena/Mala Comunicación

Al levantar un problema o buscar asistencia, por favor, sé educado en tu solicitud:

  🙂 "X no compila cuando hago Y, ¿tienes alguna sugerencia?"

  😞 "¡X no funciona! ¡Por favor, arréglalo!"

Al cerrar una Solicitud de Extracción (PR), transmite un mensaje explicativo y cordial 
que explique por qué no cumple con los requisitos para ser fusionado.

🙂 "Estoy cerrando esta PR porque esta característica no puede admitir el caso de uso X. En 
su forma propuesta, sería mejor implementarla con la herramienta Y. Gracias por trabajar en esto."

😞 "¿Por qué esto no sigue las convenciones de la API? Esto debería hacerse en otro lugar."

---

## Presentar una Contribución

### Firmar el CLA

Antes de poder presentar una contribución, debes [firmar el Acuerdo de Licencia del 
Colaborador (CLA)][cla]. 
El proyecto Kubernetes _solo_ puede aceptar una contribución si tú o tu empresa 
han firmado el CLA.

Si encuentras problemas al firmar el CLA, sigue 
las [pautas de solución de problemas del CLA].
  
### Abrir y Responder Issues

Los Issues de GitHub son el medio principal para dar seguimiento a cosas como informes de errores, 
solicitudes de mejoras o para informar otros problemas, como pruebas que fallan. 
**No** están destinados a [solicitudes de soporte de usuarios]. Para ello, consulta la [guía de solución de problemas], 
informa el problema en [Stack Overflow] o sigue en el [foro de Kubernetes].

**Referencias:**
- [Labels]
- [Comandos de Prow][comandos]

#### Crear un Issue

- Utiliza una plantilla de issue si está disponible. Utilizar la correcta plantilla
  ayudará a otros colaboradores a responder a tu issue.
  - Sigue cualquier instrucción descrita en la plantilla del issue.
- Sé descriptivo en el issue que estás planteando.
- Asigna las [labels] adecuadas. Si tienes dudas, el bot [k8s-ci-robot][prow]
  (bot [Kubernetes CI][prow]) responderá a tu issue con las etiquetas necesarias
  para un triage correcto.
- Sé selectivo al asignar issues usando [`/assign @<nombre de usuario>`][assign] o
  [`/cc @<nombre de usuario>`][cc]. Tu issue se triageará de manera más efectiva aplicando
  labels correctas en lugar de asignar más personas al issue.


#### Responder a un Issue

- Cuando estés abordando un issue, comenta en él para que otros sepan que estás trabajando en él
  y evitar trabajo duplicado.
- Cuando hayas resuelto algo por ti mismo en un momento posterior, comenta en el issue
  para que la gente sepa antes de cerrarlo.
- Incluye referencias a otros PRs o issues (o cualquier material accesible),
  por ejemplo: _"ref: #1234"_. Es útil identificar que el trabajo relacionado se ha
  abordado en otro lugar.


### Abrir un Pull Request

Las solicitudes de extracción (Pull Request o PR) son el principal medio 
para contribuir con código, documentación u
otras formas de trabajo que se almacenarán en un repositorio git.

**Referencias:**
- [Labels]
- [Comandos de Prow][comandos]
- [Proceso de Solicitud de Extracción (Pull Request)][Pull Requests]
- [Flujo de Trabajo de GitHub]


#### Crear un Pull Request

- Sigue las instrucciones de la plantilla de PR si está disponible. 
  Ayudará a quienes responden a tu PR.
- Si se trata de una [corrección trivial], como un enlace roto, un error tipográfico o de gramática, revisa el
  documento completo en busca de otros errores potenciales. No abras múltiples PR para
  pequeñas correcciones en el mismo documento.
- Hace referencia a los problemas relacionados con tu PR o a los problemas que el PR podría resolver.
- Evita realizar cambios excesivamente grandes en un solo commit. En su lugar, divide tu PR
  en varios commits pequeños y lógicos. Esto facilita que tu PR sea revisado.
- Comenta en tu propio PR cuando creas que algo necesita una explicación adicional.
- Sé selectivo al asignar tu PR con [`/assign @<nombre de usuario>`][assign].
  Asignar revisores en exceso no acelerará la revisión de tu PR.
- Si tu PR se considera un _"Trabajo en progreso"_(_"Work in progress"_), prefixa el nombre con `[WIP]`
  o utiliza el comando [`/hold`][hold]. Esto evitará que el PR se fusiona hasta que se levante el `[WIP]` o la retención.
- Si tu PR no recibe suficiente atención, publica un enlace al PR en el
  canal `#pr-reviews` en Slack para encontrar revisores adiccionales.


#### Ejemplo de Descripción de un Pull Request

```
Ref. #3064 #3097
All files owned by SIG testing were moved from `/devel` to the new folder `/devel/sig-testing`.

/sig contributor-experience
/cc @stakeholder1 @stakeholder2
/kind cleanup
/area developer-guide
/assign @approver1 @approver2 @approver3
```

Lo que contiene ese PR:
- **Línea 1** - Referencia a otros issues o PRs (#3064 #3097).
- **Línea 2** - Una breve descripción de lo que se está haciendo en el PR.
- **Línea 4** - Asignación de [SIG][SIGs] con el [comandos]
  `/sig contributor-experience`.
- **Línea 5** - Los revisores que pueden estar interesados en este problema o PR están
  especificados con el comando [`/cc`][cc].
- **Línea 6** - El comando [`/kind cleanup`][kind] agrega una [etiqueta][labels] que
  categoriza el problema o PR como relacionado con la limpieza de código,
  procesos o deuda técnica.
- **Línea 7** - El comando [`/area developer-guide`][kind] categoriza el problema o
  PR como relacionado con la guía del desarrollador.
- **Línea 8** - El comando [`/assign`][assign] asigna un aprobador al PR.
  El [k8s-ci-robot][prow] sugerirá un aprobador y se seleccionará de la
  lista de propietarios en el archivo [OWNERS]. Agregarán la
  etiqueta [`/approve`][approve] al PR después de que haya sido revisado.


#### Solucionar Problemas de un Pull Request

Después de que se propone tu PR, se ejecutan una serie de pruebas en la plataforma de CI 
de Kubernetes, [Prow]. Si alguna de las pruebas falla, el [k8s-ci-robot][prow]
responderá al PR con enlaces a las pruebas fallidas y logs disponibles.

Agregar nuevos commits a tu PR activará automáticamente la repetición de las pruebas.

Ocasionalmente puede haber problemas con la plataforma de CI de Kubernetes. Estos pueden ocurrir
por una variedad de razones, incluso si tu contribución pasa todas las pruebas locales. 
Puedes activar una repetición de las pruebas con el comando `/retest`.

Para obtener más información sobre la solución de problemas de pruebas específicas, consulta la [Guía de Pruebas].


### Labels

Kubernetes utiliza [labels] para categorizar y priorizar issues y Pull Requests.
Aplicar los labels adecuados ayudará a que tu problema o PR se priorice de manera más
efectiva.

**Referencias:**
- [Labels]
- [Comandos de Prow][Comandos]

Labels frecuentemente utilizados:
- [`/sig <nombre de SIG>`][kind] Asigna un [SIG][SIGs] a la propiedad del issue
  o PR.
- [`/area <nombre de área>`][kind] Asocia el issue o las PRs a un área específica
  [área][labels].
- [`/kind <categoría>`][kind] [Categoriza][labels] el problema o PR.

---

## Trabajar Localmente

Antes de proponer una solicitud de extracción (Pull Request), deberás realizar algún trabajo 
localmente. 
Si eres nuevo en Git, el [tutorial de Git de Atlassian] es un buen punto de partida. 
Como alternativa, el [tutorial Git Magic] de Stanford es una buena opción multilingüe.

**Referencias:**
- [Tutorial de Git de Atlassian]
- [tutorial Git Magic]
- [Flujo de trabajo de GitHub]
- [Pruebas locales]
- [Guía del desarrollador]

### Estrategia de Branch

El proyecto Kubernetes utiliza un flujo de trabajo de _"Fork y Pull"_ que es estándar en GitHub. 
En términos de Git, tu propio fork se denomina _"`origin`"_ y el repositorio Git del proyecto real se llama _"`upstream`"_. 
Para mantener tu branch personal (`origin`) actualizada con el proyecto (`upstream`), 
debes configurarla en tu copia de trabajo local.

#### Agregar Upstream

Agrega `upstream` como un repositorio remoto y configúralo para que no puedas enviar cambios a él.

```
# Reemplaza <repositorio Git upstream> con la URL del repositorio upstream
# Ejemplo:
#  https://github.com/kubernetes/kubernetes.git
#  git@github.com/kubernetes/kubernetes.git

git remote add upstream <repositorio Git upstream>
git remote set-url --push upstream no_push
```

Esto se puede verificar ejecutando `git remote -v`, 
que mostrará tus repositorios remotos configurados.

#### Mantener Tu Fork Sincronizado

Obtén todos los cambios de `upstream` y _"rebasa"_ (rebase) en tu branch local `master`. 
Esto sincronizará tu repositorio local con el proyecto `upstream`. 
Luego, envía los cambios locales a tu `remote master`.

```
git fetch upstream
git checkout master
git rebase upstream/master
git push
```

Esto es lo mínimo que debes hacer antes de crear una nueva rama para trabajar en tu función o corrección.

```
git checkout -b myfeature
```

#### Squashing Commits

El propósito principal de [squashing commits] es crear un Git limpio con 
historial legible y con log de los cambios realizados. 
Normalmente, esto se hace en la última fase de revisión de un PR. 
Si no estás seguro de si debes hacer el squashing de tus commits, es mejor optar por tener más commits y dejar que lo decidan los otros colaboradores encargados de revisar y aprobar tu PR.

Realiza un rebase interactivo para elegir qué commits deseas conservar y cuáles deseas combinar, y luego fuerza la actualización de tu branch:

```
git rebase -i HEAD~3
...
git push --force
```

**Nota**: También puedes pedir a tu revisor que añada el label `tide/merge-method-squash` a 
tu PR (esto puede ser hecho por un revisor emitiendo el comando: `/label tide/merge-method-squash`), esto permitirá que el bot se encargue de 
combinar _todos_ los commits que forman parte de esta PR y no resultará en la eliminación del 
label `LGTM` (si ya se ha aplicado) o en la repetición de las pruebas CI.

[Curso para Colaboradores]: https://www.kubernetes.dev/docs/onboarding
[Guía para Colaboradores]: /contributors/guide/README.md
[Guía del Desarrollador]: /contributors/devel/README.md
[Prow]: https://prow.k8s.io
[Tide]: http://git.k8s.io/test-infra/prow/cmd/tide/pr-authors.md
[Panel de Tide]: https://prow.k8s.io/tide
[Comandos de Bots]: https://go.k8s.io/bot-commands
[GitHub Labels]: https://go.k8s.io/github-labels
[Búsqueda de Código de Kubernetes]: https://cs.k8s.io/
[@dims]: https://github.com/dims
[Calendario]: https://calendar.google.com/calendar/embed?src=calendar%40kubernetes.io
[kubernetes-dev]: https://groups.google.com/a/kubernetes.io/g/dev
[Canales de Slack]: http://slack.k8s.io/
[Stack Overflow]: https://stackoverflow.com/questions/tagged/kubernetes
[Canal de YouTube]: https://www.youtube.com/c/KubernetesCommunity/
[Panel de Triage]: https://go.k8s.io/triage
[Test Grid]: https://testgrid.k8s.io
[Estadísticas de Desarrolladores]: https://k8s.devstats.cncf.io
[Código de Conducta]: /code-of-conduct.md
[Solicitudes de Soporte de Usuarios]: /contributors/guide/issue-triage.md#determine-if-its-a-support-request
[guía de solución de problemas]: https://kubernetes.io/docs/tasks/debug-application-cluster/troubleshooting/
[Foro de Kubernetes]: https://discuss.kubernetes.io/
[Pull Requests]: /contributors/guide/pull-requests.md
[Flujo de Trabajo de GitHub]: /contributors/guide/github-workflow.md
[Prow]: https://git.k8s.io/test-infra/prow#prow
[CLA]: /CLA.md#how-do-i-sign
[pautas de solución de problemas del CLA]: /CLA.md#troubleshooting
[Comandos]: https://prow.k8s.io/command-help
[kind]: https://prow.k8s.io/command-help#kind
[CC]: https://prow.k8s.io/command-help#cc
[hold]: https://prow.k8s.io/command-help#hold
[assign]: https://prow.k8s.io/command-help#assign
[SIGs]: /sig-list.md
[Guía de Pruebas]: /contributors/devel/sig-testing/testing.md
[Labels]: https://git.k8s.io/test-infra/label_sync/labels.md
[corrección trivial]: /contributors/guide/pull-requests.md#10-trivial-edits
[Flujo de Trabajo de GitHub]: /contributors/guide/github-workflow.md#3-branch
[Squashing Commits]: /contributors/guide/pull-requests.md#6-squashing-and-commit-titles
[OWNERS]: /contributors/guide/owners.md
[Pruebas Locales]: /contributors/devel/sig-testing/testing.md
[Tutorial de Git de Atlassian]: https://www.atlassian.com/git/tutorials
[tutorial Git Magic]: http://www-cs-students.stanford.edu/~blynn/gitmagic/
[Guía de Seguridad y Divulgación]: https://kubernetes.io/docs/reference/issues-security/security/
[approve]: https://prow.k8s.io/command-help#approve
[Equipo de Administración de GitHub]: /github-management#github-administration-team
[Publicación de Parches de Kubernetes]: https://kubernetes.io/releases/patch-releases/

