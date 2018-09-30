@extends('layout')

@section('content')
<script>
var varCount = 1;
var varCount2 = 1;
 
$(function () {
    $('#new').on('click', function(){
        varCount++;
        $node = '<div id="in'+varCount+'" class="input-group mb-3"><div class="input-group-prepend"><span class="input-group-text" id="basic-addon1">'+varCount+'</span></div><textarea rows="1" name="input[]" class="form-control" aria-label="With textarea"></textarea></div>';
        $node2 = '<div id="out'+varCount+'" class="input-group mb-3"><div class="input-group-prepend"><span class="input-group-text" id="basic-addon1">'+varCount+'</span></div><textarea rows="1" name="output[]" class="form-control" aria-label="With textarea"></textarea></div>';
        $('#in').parent().append($node);
        $('#out').parent().append($node2);
    });

    $('#new_ex').on('click', function(){
        varCount2++;
        $node = '<div id="ex'+varCount2+'" class="input-group mb-3"><div class="input-group-prepend"><span class="input-group-text" id="basic-addon1">'+varCount2+'</span></div><textarea rows="3" name="example[]" class="form-control" aria-label="With textarea"></textarea></div>';
        $('#ex').parent().append($node);
    });

    $('#del').on('click', function(){
        if (varCount>1){
            $('#in'+varCount).remove();
            $('#out'+varCount).remove();
            varCount--;
        }
    });
    $('#del_ex').on('click', function(){
        if (varCount2>1){
            $('#ex'+varCount2).remove();
            varCount2--;
        }
    });
});

function tab(obj){
    if (event.keyCode == 9)
    {
        obj.value = obj.value + "    "; 
        event.returnValue = false;
    }
}

</script>
<section class="container">
{{Auth::user()->class}}

<form aciton="{{ url('questionpool/create')}}" method="post">
    {{ csrf_field() }}
    @if (isset($data->id))
        <input type="hidden" name="num" value="{{$data->id + 1}}">
    @else
        <input type="hidden" name="num" value="1">
    @endif
    
    <table class="table table-hover">
    <tr>
        <td>Title
        <input type="text" placeholder="title" name="title" class="form-control">
        </td>
    </tr>
    <tr>
        <td>Content
        <textarea cols="50" rows="5" name="content" class="form-control" ></textarea>
        </td>
    </tr>
    <tr>
        <td>Code Function
        <textarea cols="50" rows="5" name="code_function" class="form-control" onkeydown="tab(this)">class Solution:</textarea>
        </td>
    </tr>
    <tr>
        <td>Example
        <button id="new_ex" type="button" class="btn btn-outline-secondary btn-sm">New</button>
        <button id="del_ex" type="button" class="btn btn-outline-danger btn-sm">Delete</button>
        <div id="ex" class="input-group mb-3">
            <div class="input-group-prepend">
                <span class="input-group-text" id="basic-addon1">1</span>
            </div>
            <textarea rows="3" name="example[]" class="form-control" aria-label="With textarea"></textarea>
        </div>
        </td>
    </tr>
    <tr>
        <td>Testing<br>Input
        <button id="new" type="button" class="btn btn-outline-secondary btn-sm">New</button>
        <button id="del" type="button" class="btn btn-outline-danger btn-sm">Delete</button>
        Type
        <select name="inputtype">
            <option value="string">String</option>
            <option value="int">Int</option>
        </select>
        <div id="in" class="input-group mb-3">
            <div class="input-group-prepend">
                <span class="input-group-text" id="basic-addon1">1</span>
            </div>
            <textarea rows="1" name="input[]" class="form-control" aria-label="With textarea"></textarea>
        </div>
        </td>
    </tr>
    <tr>
        <td>Testing<br>Output
        <div id="out" class="input-group mb-3">
            <div class="input-group-prepend">
                <span class="input-group-text" id="basic-addon1">1</span>
            </div>
            <textarea rows="1" name="output[]" class="form-control" aria-label="With textarea"></textarea>
        </div>
        </td>
    </tr>
    </table>
    <input type="submit" class="btn btn-dark btn-lg">
</form>
</section>
@stop