import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateEditResourceComponent } from './create-edit-resource.component';

describe('CreateEditResourceComponent', () => {
  let component: CreateEditResourceComponent;
  let fixture: ComponentFixture<CreateEditResourceComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CreateEditResourceComponent]
    });
    fixture = TestBed.createComponent(CreateEditResourceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
