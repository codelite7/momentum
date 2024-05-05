import { TestBed } from '@angular/core/testing';
import { CanActivateChildFn } from '@angular/router';

import { canActivateAppGuard } from './can-activate-app.guard';

describe('canActivateAppGuard', () => {
  const executeGuard: CanActivateChildFn = (...guardParameters) => 
      TestBed.runInInjectionContext(() => canActivateAppGuard(...guardParameters));

  beforeEach(() => {
    TestBed.configureTestingModule({});
  });

  it('should be created', () => {
    expect(executeGuard).toBeTruthy();
  });
});
