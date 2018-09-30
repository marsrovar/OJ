<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Questionpool;
use App\Questionpools_example;
use App\Questionpools_testing;

class QuestionPoolController extends Controller
{

    public function __construct()
    {
        $this->middleware('auth');
    }

    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request)
    {
        // $request->session()->put('key', 'a');
        // $value = $request->session()->get('key');
        // $request->session()->forget('key');
        // $request->session()->flush();
        $data = Questionpool::all();
        return view('questionpool.index', compact('data'));
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        $data = Questionpool::orderBy('id', 'desc')->first();
        return view('questionpool.create', compact('data'));
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        $pattern='\t';
        $code_fun= preg_replace("/($pattern)/i",'    ',$request->code_function);
        Questionpool::create([
            'num' => $request->num,
            'title' => $request->title,
            'content' => $request->content,
            'code_function' => $code_fun,
            'inputtype' => $request->inputtype,
        ]);
        $qpid = Questionpool::where('num', '=', $request->num)->pluck('id');
        $count = count($request->input);
        $count2 = count($request->example);
        for ($i = 0; $i < $count; $i++) {
            Questionpools_testing::create([
                    'qpid' => $qpid[0],
                    'input' => $request->input[$i],
                    'output' => $request->output[$i],
                    ]);
        }
        for ($i = 0; $i < $count2; $i++) {
            Questionpools_example::create([
                    'qpid' => $qpid[0],
                    'example' => $request->example[$i],
                    ]);
        }
        return redirect('questionpool');
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        $dataQuestion = Questionpool::where('id', '=', $id)->get();
        $dataExample = Questionpools_example::where('qpid', '=', $id)->get();
        return view('questionpool.show', compact('dataQuestion','dataExample'));
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $id)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        //
    }
}