@extends('layout')

@section('content')
<section class="container">
<a href="{{ url('questionpool') }}" role="btn" class="btn btn-primary">Back</a>
    <table class="table table-hover">
    @foreach ($dataQuestion as $var)
    <tr>
        <td>Title
        <textarea cols="50" rows="1" name="title" class="form-control" disabled>{{ $var->title }}</textarea>
        </td>
    </tr>
    <tr>
        <td>Content
        <textarea cols="50" rows="5" name="content" class="form-control" disabled>{{ $var->content }}</textarea>
        </td>
    </tr>
    <tr>
        <td>Code Function
        <textarea cols="50" rows="5" name="code_function" class="form-control" disabled>{{ $var->code_function }}</textarea>
        </td>
    </tr>
    @endforeach
    <tr>
        <td>Example
        @foreach ($dataExample as $var)
        <div  class="input-group mb-3">
            <div class="input-group-prepend">
                <span class="input-group-text" id="basic-addon1">{{ $loop->index +1 }}</span>
            </div>
            <textarea cols="50" rows="3" name="content" class="form-control" disabled>{{ $var->example }}</textarea>
        </div>
        @endforeach
        </td>
    </tr>
    </table>
</section>
@stop