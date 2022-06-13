# SIKerjasama - Frontend _(unhan/sikerjasama/frontend)_

> _Frontend_ Sistem Pengelolaan Data Kerjasama (SIKerjasama).

## Pipeline Status

- [`develop`](https://kerma.d.pusilkom.com/) [![pipeline status](https://gitlab.pusilkom.com/unhan/sikerjasama/frontend/badges/develop/pipeline.svg)](https://gitlab.pusilkom.com/unhan/sikerjasama/frontend/commits/develop)
- [`staging`](https://kerma.s.pusilkom.com/) [![pipeline status](https://gitlab.pusilkom.com/unhan/sikerjasama/frontend/badges/testing/pipeline.svg)](https://gitlab.pusilkom.com/unhan/sikerjasama/frontend/commits/testing)
- [`production`](https://kerma.idu.ac.id/) [![pipeline status](https://gitlab.pusilkom.com/unhan/sikerjasama/frontend/badges/production/pipeline.svg)](https://gitlab.pusilkom.com/unhan/sikerjasama/frontend/commits/production)

## Project Setup

```bash
npm install
```

Prepare a `.env` file in the project directory before attempting to build
the application locally. An example of `.env` file can be found at
[`.env-example`](./.env-example) file. The list of variables and their purpose
are as follows:

- `VUE_APP_BACKEND_URL` contains the URL to the backend API. Defaults to
  `http://localhost:8001` if not defined.
- `VUE_APP_SSO_URL` contains the URL to the SSO. Defaults to
  `https://sso.d.pusilkom.com` if not defined.
- `VUE_APP_USERMANAGEMENT_URL` contains the URL to the user management API.
  Defaults to `https://user.d.pusilkom.com` if not defined.
- `VUE_APP_CHAT_URL` contains the URL to the chat (messaging) system.
  Defaults to `https://chat.d.pusilkom.com` if not defined.
- `VUE_APP_LOGGING_URL` contains the URL to the traffic logging system.
  Defaults to `https://unhan-graylog.pusilkom.com` if not defined.
- `VUE_APP_SITEKEY` contains ReCaptcha site key. Defaults to empty string if
  not defined.
- `VUE_APP_CLIENT_ID` contains the client ID. Defaults to empty string if not
  defined.
- `VUE_APP_CLIENT_SECRET` contains the client secret. Defauls to empty string
  if not defined.
- `VUE_APP_SSO_APIKEY` contains the API key to access the SSO system.
  Defaults to empty string if not defined.

### Compiles and hot-reloads for development

```bash
npm run serve
```

### Compiles and minifies for production

```bash
npm run build
```

### Lints and fixes files

```bash
npm run lint
```

### Customize configuration

See [Configuration Reference](https://cli.vuejs.org/config/).
