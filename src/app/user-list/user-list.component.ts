import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css'],
})
export class UserListComponent implements OnInit {
  users: any[] = [];
  searchQuery: string = '';
  currentPage: number = 1;
  totalPages: number = 1;

  constructor(private http: HttpClient, private router: Router) {}

  ngOnInit(): void {
    this.fetchUsers(this.currentPage);
  }

  fetchUsers(page: number): void {
    const apiUrl = `https://reqres.in/api/users?page=${page}`;
    this.http.get(apiUrl).subscribe((data: any) => {
      this.users = data.data;
      this.currentPage = data.page;
      this.totalPages = data.total_pages;
    });
  }

  navigateToUserDetails(userId: number): void {
    this.router.navigate(['user', userId]);
  }

  prevPage(): void {
    if (this.currentPage > 1) {
      this.fetchUsers(this.currentPage - 1);
    }
  }

  nextPage(): void {
    if (this.currentPage < this.totalPages) {
      this.fetchUsers(this.currentPage + 1);
    }
  }
}
