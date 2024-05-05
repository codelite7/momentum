import { CanActivateFn, Router } from '@angular/router'
import { inject } from '@angular/core'
import { AuthService } from '../auth.service'

export const canActivateSignOutGuard: CanActivateFn = async (route, state) => {
  const router = inject(Router);
  let sessionExists = await inject(AuthService).sessionExists()
  if (sessionExists) {
    return true
  } else {
    return router.parseUrl('auth/signin')
  }
};
