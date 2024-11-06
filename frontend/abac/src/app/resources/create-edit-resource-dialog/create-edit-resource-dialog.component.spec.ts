import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateEditResourceDialogComponent } from './create-edit-resource-dialog.component';

describe('CreateEditResourceDialogComponent', () => {
  let component: CreateEditResourceDialogComponent;
  let fixture: ComponentFixture<CreateEditResourceDialogComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CreateEditResourceDialogComponent]
    });
    fixture = TestBed.createComponent(CreateEditResourceDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
