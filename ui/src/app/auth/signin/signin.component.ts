import { Component, inject } from '@angular/core'
import { FloatLabelModule } from 'primeng/floatlabel'
import { FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms'
import { NgOptimizedImage } from '@angular/common'
import { InputTextModule } from 'primeng/inputtext'
import { PasswordModule } from 'primeng/password'
import { ButtonModule } from 'primeng/button'
import { emailPasswordSignIn } from 'supertokens-web-js/lib/build/recipe/thirdpartyemailpassword'
import { Router, RouterLink } from '@angular/router'
import { ToastService } from '../../services/toast.service'

@Component({
  selector: 'app-signin',
  standalone: true,
  imports: [
    FloatLabelModule,
    ReactiveFormsModule,
    NgOptimizedImage,
    InputTextModule,
    FormsModule,
    PasswordModule,
    ButtonModule,
    RouterLink
  ],
  templateUrl: './signin.component.html',
  styleUrl: './signin.component.css'
})
export class SigninComponent {
  toastService: ToastService = inject(ToastService)
  router: Router = inject(Router)
  signInForm: FormGroup = new FormGroup<any>({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required])
  })
  loading: boolean = false;

  async signIn() {
    try {
      this.loading = true;
      let response = await emailPasswordSignIn({
        formFields: [{
          id: "email",
          value: this.signInForm.value['email'],
        }, {
          id: "password",
          value: this.signInForm.value['password']
        }]
      })

      if (response.status === "FIELD_ERROR") {
        response.formFields.forEach(formField => {
          if (formField.id === "email") {
            // Email validation failed (for example incorrect email syntax).
            this.toastService.error(formField.error)
          }
        })
      } else if (response.status === "WRONG_CREDENTIALS_ERROR") {
        this.toastService.error("Email password combination is incorrect.")
      } else if (response.status === "SIGN_IN_NOT_ALLOWED") {
        // the reason string is a user friendly message about what went wrong. It can also contain a support code which
        // users can tell you so you know why their sign in was not allowed.
        this.toastService.error(response.reason)
      } else {
        // sign in successful. The session tokens are automatically handled by the frontend SDK.
        this.router.navigate(['/'])
      }
    } catch (err: any) {
      if (err.isSuperTokensGeneralError === true) {
        // this may be a custom error message sent from the API by you.
        this.toastService.error(err.message);
      } else {
        this.toastService.error(err);
      }
    } finally {
      this.loading = false;
    }
  }
}
