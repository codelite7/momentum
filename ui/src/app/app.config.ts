import { ApplicationConfig } from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { provideAnimations } from '@angular/platform-browser/animations'
import { MessageService } from 'primeng/api'

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(routes),
    provideAnimations(),
    MessageService, // the toast service needs the message service, this instantiates the message services so the injector works in the toast service
  ]
};
