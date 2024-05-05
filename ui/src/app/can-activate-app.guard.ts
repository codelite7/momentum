import { CanActivateChildFn, Router, UrlTree } from '@angular/router'
import { AuthService } from './auth.service'
import { inject } from '@angular/core'

export const canActivateAppGuard: CanActivateChildFn = async (childRoute, state) => {
  const router = inject(Router);
  let sessionExists = await inject(AuthService).sessionExists()
  if (sessionExists) {
    return true
  } else {
    try {
      let tree = router.parseUrl('auth/signin')
      return tree
    } catch (e) {
      console.error(e)
      return false
    }
  }
};
