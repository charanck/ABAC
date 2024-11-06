import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable, tap } from 'rxjs';

const BASE_URL = 'http://localhost:3000';
@Injectable({
  providedIn: 'root',
})
export class ResourceService {
  constructor(private http: HttpClient) {}

  listResources(
    pageNumber: number,
    pageSize: number
  ): Observable<{
    data: Resource[];
    pagingMetadata: PagingMetadata;
  }> {
    return this.http
      .get<{
        data: Resource[];
        pagingMetadata: PagingMetadata;
      }>(`${BASE_URL}/resources`, {
        params: {
          pageNumber: pageNumber,
          pageSize: pageSize,
        },
      })
      .pipe(
        map((response) => {
          for (let i = 0; i < response.data.length; i++) {
            response.data[i].deleted = new Date(response.data[i].deleted || 0);
            response.data[i].created = new Date(response.data[i].created || 0);
            response.data[i].updated = new Date(response.data[i].updated || 0);
          }
          return {
            data: response.data,
            pagingMetadata: {
              total: response.pagingMetadata.total,
            },
          };
        })
      );
  }

  createResource(resource: Resource): Observable<void> {
    return this.http
      .post<any>(`${BASE_URL}/resources`, resource)
      .pipe(
        tap((res) => {
          console.log('Res: ', res);
        })
      );
  }
}

export interface Resource {
  description: string;
  id?: string;
  name: string;
  ownerId: string;
  policyId: string;
  updated?: Date;
  created?: Date;
  deleted?: Date;
}

export interface PagingMetadata {
  total: number;
}
