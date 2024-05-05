import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { AppComponent } from './app/app.component';
import SuperTokens from 'supertokens-web-js';
import Session from 'supertokens-web-js/recipe/session';
import ThirdPartyEmailPassword from 'supertokens-web-js/recipe/thirdpartyemailpassword'

SuperTokens.init({
  appInfo: {
    apiDomain: "http://localhost:3000",
    apiBasePath: "/auth",
    appName: "momentum",
  },
  recipeList: [
    Session.init(),
    ThirdPartyEmailPassword.init(),
  ],
});

bootstrapApplication(AppComponent, appConfig)
  .catch((err) => console.error(err));
