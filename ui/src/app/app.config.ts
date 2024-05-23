import { ApplicationConfig, importProvidersFrom } from '@angular/core'
import { provideRouter, withComponentInputBinding, withRouterConfig } from '@angular/router'

import { routes } from './app.routes';
import { provideAnimations } from '@angular/platform-browser/animations'
import { MessageService } from 'primeng/api';
import { provideHttpClient } from '@angular/common/http';
import { graphqlProvider } from './graphql.provider'
import { GraphqlService } from './services/graphql.service'
import { MarkdownModule } from 'ngx-markdown'
import { NgxCutModule } from 'ngx-cut'

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(routes, withComponentInputBinding(), withRouterConfig({
      paramsInheritanceStrategy: 'always'
    })),
    provideAnimations(),
    MessageService, // the toast service needs the message service, this instantiates the message services so the injector works in the toast service
    GraphqlService, // individual object services inject the graphql service so it gets instantiated here
    provideHttpClient(),
    graphqlProvider,
    importProvidersFrom(
      MarkdownModule.forRoot(),
      NgxCutModule.forRoot({
        size: 1,
        breakpoints: { sm: 300, md: 400, lg: 500, xl: 600 },
        responsiveSizes: {
          xs: { xs: 1, sm: 2, md: 3, lg: 4, xl: 5 },
          sm: { xs: 2, sm: 3, md: 4, lg: 5, xl: 6 },
          md: { xs: 3, sm: 4, md: 5, lg: 6, xl: 7 },
          lg: { xs: 4, sm: 5, md: 6, lg: 7, xl: 8 },
          xl: { xs: 5, sm: 6, md: 7, lg: 8, xl: 9 }
        }
      })
    )
  ]
};
