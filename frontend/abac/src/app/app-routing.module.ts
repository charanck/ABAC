import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ResourcesComponent } from './resources/resources.component';
import { CreateEditResourceDialogComponent } from './resources/create-edit-resource-dialog/create-edit-resource-dialog.component';

const routes: Routes = [
  {
    path: "resources",
    component: ResourcesComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
