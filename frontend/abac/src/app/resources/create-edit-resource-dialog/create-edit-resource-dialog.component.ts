import { Component, inject } from '@angular/core';
import { Resource, ResourceService } from '../resource.service';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { catchError, EMPTY } from 'rxjs';
import { SnackbarService } from 'src/app/util/snackbar.service';

enum FormType {
  edit = 1,
  create = 2,
}
@Component({
  selector: 'app-create-edit-resource-dialog',
  templateUrl: './create-edit-resource-dialog.component.html',
  styleUrls: ['./create-edit-resource-dialog.component.scss'],
})
export class CreateEditResourceDialogComponent {
  protected readonly FormType = FormType;
  private readonly dialogRef = inject(
    MatDialogRef<CreateEditResourceDialogComponent>
  );
  private readonly data = inject<Resource>(MAT_DIALOG_DATA);
  private readonly resourceService = inject(ResourceService);
  private readonly snackbarService = inject(SnackbarService);
  protected type: FormType = FormType.create;
  protected form: FormGroup = new FormGroup({
    name: new FormControl(this.data?.name || '', [Validators.required]),
    description: new FormControl(this.data?.description || '', [
      Validators.required,
    ]),
    ownerId: new FormControl(this.data?.ownerId || '', [Validators.required]),
    policyId: new FormControl(this.data?.policyId || '', [Validators.required]),
  });

  constructor() {
    this.type = this.data?.id ? FormType.edit : FormType.create;
  }

  create() {
    const resource: Resource = {
      description: this.form.get('description')?.value,
      name: this.form.get('name')?.value,
      ownerId: this.form.get('ownerId')?.value,
      policyId: this.form.get('policyId')?.value,
    };
    this.resourceService
      .createResource(resource)
      .pipe(
        catchError((err) => {
        console.log(err);
          this.snackbarService.success('Resource Creation failed');
          return EMPTY;
        })
      )
      .subscribe(() => {
        this.snackbarService.success('Resource created successfully');
      });
  }

  update() {}

  close() {
    this.dialogRef.close();
  }
}
