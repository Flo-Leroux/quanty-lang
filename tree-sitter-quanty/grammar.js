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

    rules: {
        // TODO: add the actual grammar rules
        program: $ => repeat($._fragment),

        _fragment: $ => choice(
            $.component_definition
        ),

        component_definition: $ => seq(
            'component',
            field('name', $.identifier),
            '{',

            '}'
        ),

        identifier: ($) => /[_A-Za-z][_0-9A-Za-z]*/,
    }
});
