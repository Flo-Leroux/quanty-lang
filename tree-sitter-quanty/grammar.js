/**
 * @file Quanty grammar for tree-sitter
 * @author Flo Leroux
 */

/* eslint-disable arrow-parens */
/* eslint-disable camelcase */
/* eslint-disable-next-line spaced-comment */
/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
    name: 'quanty',

    externals: ($) => [
        $._newline,
        $._indent,
        $._dedent,
    ],

    rules: {
        program: $ => repeat($._fragment),

        _fragment: $ => choice(
            $.component_declaration,
        ),

        component_declaration: $ => seq(
            'component',
            field('name', $.identifier),
            $._indent,
            field('body', repeat1($.tag)),
            $._dedent,
        ),

        tag: $ => seq(
            field('name', $.identifier),
            field('children', optional($.children))
        ),

        children: ($) => seq(
            $._indent,
            repeat1($.tag),
            $._dedent
        ),

        identifier: ($) => /[_A-Za-z][_0-9A-Za-z]*/,
    }
});
