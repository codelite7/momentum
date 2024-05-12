# Ui

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 17.3.6.

## Design principals
### API Interaction
All api interactions go through services, there is one service for each graphql type. There is a base graphl service that makes the requests.
### Emitters
Mutations don't return values, instead they emit the id of the thing that was mutated. This is because various components have different requirements 
for fields and child objects that they care about. Each component that cares about mutations listens to the emitter it cares about and fetches the data
it needs to render. This reduces bugs and brittle code by allowing each component to only care about the fields it needs without tyring to modify
objects to make them look the way the component wants.

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The application will automatically reload if you change any of the source files.

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via a platform of your choice. To use this command, you need to first add a package that implements end-to-end testing capabilities.

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.
