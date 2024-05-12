import { Component, inject } from '@angular/core'
import { ButtonModule } from 'primeng/button'
import { FloatLabelModule } from 'primeng/floatlabel'
import { InputTextModule } from 'primeng/inputtext'
import { NgOptimizedImage } from '@angular/common'
import { PaginatorModule } from 'primeng/paginator'
import { PasswordModule } from 'primeng/password'
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms'
import { emailPasswordSignUp } from 'supertokens-web-js/recipe/thirdpartyemailpassword'
import { ToastService } from '../../services/toast.service'
import { Router, RouterLink } from '@angular/router'

@Component({
  selector: 'app-signup',
  standalone: true,
  imports: [
    ButtonModule,
    FloatLabelModule,
    InputTextModule,
    NgOptimizedImage,
    PaginatorModule,
    PasswordModule,
    ReactiveFormsModule,
    RouterLink
  ],
  templateUrl: './signup.component.html',
  styleUrl: './signup.component.css'
})
export class SignupComponent {
  toastService: ToastService = inject(ToastService)
  router: Router = inject(Router)
  signUpForm: FormGroup = new FormGroup<any>({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required])
  })
  loading: boolean = false;
  async signUp() {
    try {
      let response = await emailPasswordSignUp({
        formFields: [{
          id: "email",
          value: this.signUpForm.value.email
        }, {
          id: "password",
          value: this.signUpForm.value.password
        }]
      })

      if (response.status === "FIELD_ERROR") {
        // one of the input formFields failed validaiton
        response.formFields.forEach(formField => {
          if (formField.id === "email") {
            // Email validation failed (for example incorrect email syntax),
            // or the email is not unique.
            this.toastService.error(formField.error)
          } else if (formField.id === "password") {
            // Password validation failed.
            // Maybe it didn't match the password strength
            this.toastService.error(formField.error)
          }
        })
      } else if (response.status === "SIGN_UP_NOT_ALLOWED") {
        // the reason string is a user friendly message
        // about what went wrong. It can also contain a support code which users
        // can tell you so you know why their sign up was not allowed.
        this.toastService.error(response.reason)
      } else {
        // sign up successful. The session tokens are automatically handled by
        // the frontend SDK.
        this.router.navigate(['/'])
      }
    } catch (err: any) {
      if (err.isSuperTokensGeneralError === true) {
        // this may be a custom error message sent from the API by you.
        this.toastService.error(err.message);
      } else {
        this.toastService.error("Oops! Something went wrong.");
      }
    }
  }
}
