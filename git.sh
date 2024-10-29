# konfigurasi git
git config --global user.name "Your Name"   # Menyetel nama pengguna
git config --global user.email "your.email@example.com"   # Menyetel email pengguna
git config --list   # Melihat semua konfigurasi

# Inisialisasi dan Kloning
git init   # Menginisialisasi repositori Git baru
git clone <url>   # Mengkloning repositori dari URL

# Menambah dan Menghapus File
git add <file>   # Menambah file ke staging area
git add .   # Menambah semua file yang telah diubah
git rm <file>   # Menghapus file dari repositori dan staging area
git rm --cached <file>   # Menghapus file hanya dari staging area

# Commit perubahan
git commit -m "Pesan commit"   # Membuat commit dengan pesan
git commit -a -m "Pesan commit"   # Commit semua perubahan yang ada

#Branching
git branch   # Melihat daftar branch
git branch <nama-branch>   # Membuat branch baru
git checkout <nama-branch>   # Beralih ke branch lain
git checkout -b <nama-branch>   # Membuat dan langsung beralih ke branch baru
git merge <nama-branch>   # Menggabungkan branch ke branch aktif
git branch -d <nama-branch>   # Menghapus branch

# pembaharuan dan sinkronisasi
git pull   # Menarik perubahan dari remote ke lokal
git push   # Mendorong commit ke repositori remote
git push origin <nama-branch>   # Mendorong branch tertentu ke remote

# Melihat Status dan Log
git status   # Melihat status file di repositori
git log   # Melihat riwayat commit
git log --oneline   # Melihat riwayat commit dalam satu baris per commit
git diff   # Melihat perubahan yang belum di-staging
git diff <commit> <commit>   # Melihat perbedaan antar-commit

#Tagging
git tag <tag-name>   # Membuat tag
git tag   # Menampilkan semua tag
git push origin <tag-name>   # Mendorong tag ke remote

# Stashing
git stash   # Menyimpan perubahan sementara
git stash list   # Melihat daftar stash
git stash apply   # Menerapkan perubahan stash terakhir
git stash drop   # Menghapus stash terakhir


#Undo Changes
git reset <file>   # Menghapus file dari staging area
git reset --hard <commit>   # Mengembalikan repositori ke commit tertentu
git revert <commit>   # Membuat commit baru yang membalikkan perubahan commit tertentu


