import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { AppComponent } from './app/app.component';
import SuperTokens from 'supertokens-web-js';
import Session from 'supertokens-web-js/recipe/session';
import ThirdPartyEmailPassword from 'supertokens-web-js/recipe/thirdpartyemailpassword'
import { environment } from './environments/environment'

SuperTokens.init({
  appInfo: {
    apiDomain: environment.supertokensApiDomain,
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
