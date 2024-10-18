import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ResourcesComponent } from './resources/resources.component';
import { CreateEditResourceComponent } from './resources/create-edit-resource/create-edit-resource.component';

const routes: Routes = [
  {
    path: "resources/:resourceId",
    component: CreateEditResourceComponent
  },
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
