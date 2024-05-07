import { ApplicationConfig } from '@angular/core';
import { provideRouter, withComponentInputBinding } from '@angular/router'

import { routes } from './app.routes';
import { provideAnimations } from '@angular/platform-browser/animations'
import { MessageService } from 'primeng/api';
import { provideHttpClient } from '@angular/common/http';
import { graphqlProvider } from './graphql.provider'
import { GraphqlService } from './services/graphql.service'

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(routes, withComponentInputBinding()),
    provideAnimations(),
    MessageService, // the toast service needs the message service, this instantiates the message services so the injector works in the toast service
    GraphqlService, // individual object services inject the graphql service so it gets instantiated here
    provideHttpClient(),
    graphqlProvider,
  ]
};
