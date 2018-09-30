<!doctype html>
<html lang="{{ app()->getLocale() }}">
    <head>
        <meta charset="utf-8">
        
        <title>Laravel</title>
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css">
        <script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>
        <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.min.js"></script>
    </head>
    <body>
    <div class="d-flex flex-column flex-md-row align-items-center p-3 px-md-4 mb-3 bg-white border-bottom shadow-sm">
        <h5 class="my-0 mr-md-auto font-weight-normal">
        <a href="{{ url('/') }}"">
            Online Judge
        </a>
        </h5>
        <nav class="my-2 my-md-0 mr-md-3">
        @if(isset(Auth::user()->email))
            <a class="p-2 text-dark" href="#">Exam</a>
            <a class="p-2 text-dark" href="{{ url('/questionpool') }}">QuestionPool</a>
        @if(Auth::user()->class == 0)
            <a class="p-2 text-dark" href="#">Users</a>
        @endif

        @endif
        </nav>
        @if(isset(Auth::user()->email))
            <a class="btn btn-outline-primary" href="{{ url('/logout') }}"">Logout</a>
        @else
            <a class="btn btn-outline-primary" href="{{ url('/register') }}"">Sign up</a>
            <a class="btn btn-outline-primary" href="{{ url('/login') }}"">Login</a>
        @endif
    </div>
        @yield('content')
    </body>
</html>