import { Component, inject } from '@angular/core'
import { AuthService } from '../../auth.service'
import { Router } from '@angular/router'

@Component({
  selector: 'app-signout',
  standalone: true,
  imports: [],
  templateUrl: './signout.component.html',
  styleUrl: './signout.component.css',
})
export class SignoutComponent {
  authService: AuthService = inject(AuthService)
  router: Router = inject(Router);
  async ngOnInit() {
    try {
      await this.authService.signOut()
    } finally {
      this.router.navigateByUrl('auth/signin')
    }
  }
}
