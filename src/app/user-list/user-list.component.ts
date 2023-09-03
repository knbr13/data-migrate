import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../core/services/user-service.service';

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
  cachedUsers: any[] = [];
  errorMessage: string = '';
  isLoading: boolean = false;

  constructor(private userService: UserService, private router: Router) {}

  ngOnInit(): void {
    this.fetchUsers(this.currentPage);
  }

  fetchUsers(page: number): void {
    if (this.cachedUsers[page]) {
      this.users = this.cachedUsers[page];
      this.currentPage = page;
    } else {
      this.isLoading = true;
      this.userService.getUsers(page).subscribe({
        next: (data: any) => {
          this.users = data.data;
          this.currentPage = data.page;
          this.totalPages = data.total_pages;
          this.cachedUsers[page] = data.data;
        },
        error: (error) => {
          this.errorMessage = 'Error fetching users.';
          this.users = [];
        },
        complete: () => {
          this.isLoading = false;
        },
      });
    }
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

  onSearchInputChange(): void {
    const userId = +this.searchQuery;
    if (Number(userId)) {
      this.errorMessage = '';
      this.fetchUserById(userId);
    } else {
      this.errorMessage = '';
      if (this.searchQuery === '') {
        this.fetchUsers(1);
      } else {
        this.errorMessage = 'Invalid user ID. Please enter a number.';
        this.users = [];
      }
    }
  }

  fetchUserById(userId: number): void {
    this.userService.getUserById(userId).subscribe({
      next: (data: any) => {
        this.users = [data.data];
      },
      error: (error) => {
        this.errorMessage = 'User not found.';
        this.users = [];
      },
    });
  }
}
