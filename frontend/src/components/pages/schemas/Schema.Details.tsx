/**
 * Copyright 2022 Redpanda Data, Inc.
 *
 * Use of this software is governed by the Business Source License
 * included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
 *
 * As of the Change Date specified in that file, in accordance with
 * the Business Source License, use of this software will be governed
 * by the Apache License, Version 2.0
 */

import React from 'react';
import { message, Select, Table } from 'antd';
import { observer } from 'mobx-react';
import { appGlobal } from '../../../state/appGlobal';
import { api } from '../../../state/backendApi';
import { PageComponent, PageInitHelper } from '../Page';
import { DefaultSkeleton, Label, OptionGroup, toSafeString } from '../../../utils/tsxUtils';
import { KowlJsonView } from '../../misc/KowlJsonView';
import { uiSettings } from '../../../state/ui';
import { NoClipboardPopover } from '../../misc/NoClipboardPopover';
import { isClipboardAvailable } from '../../../utils/featureDetection';
import Section from '../../misc/Section';
import PageContent from '../../misc/PageContent';
import { makeObservable, observable } from 'mobx';
import { editQuery } from '../../../utils/queryHelper';
import { Button, Flex, Icon, Tag, Tooltip } from '@redpanda-data/ui';
import { AiOutlineCopy } from 'react-icons/ai';
import { Statistic } from '../../misc/Statistic';


@observer
class SchemaDetailsView extends PageComponent<{ subjectName: string }> {
    subjectNameRaw: string;
    subjectNameEncoded: string;

    @observable version = 'latest' as 'latest' | number;

    initPage(p: PageInitHelper): void {
        const subjectNameRaw = decodeURIComponent(this.props.subjectName);
        const subjectNameEncoded = encodeURIComponent(subjectNameRaw);

        const version = getVersionFromQuery();
        editQuery(x => {
            console.log('version', { fromQuery: version, current: x.version, currentQuery: window.location.search });
            x.version = String(this.version);
        });

        p.title = subjectNameRaw;
        p.addBreadcrumb('Schema Registry', '/schema-registry');
        p.addBreadcrumb(subjectNameRaw, `/schema-registry/${subjectNameEncoded}?version=${version}`);
        this.refreshData(false);
        appGlobal.onRefresh = () => this.refreshData(true);
    }

    constructor(p: any) {
        super(p);
        this.subjectNameRaw = decodeURIComponent(this.props.subjectName);
        this.subjectNameEncoded = encodeURIComponent(this.subjectNameRaw);
        makeObservable(this);
    }

    refreshData(force?: boolean) {
        api.refreshSchemaConfig(force);
        api.refreshSchemaMode(force);
        api.refreshSchemaSubjects(force);
        api.refreshSchemaTypes(force);

        const encoded = decodeURIComponent(this.props.subjectName);
        api.refreshSchemaDetails(encoded, 'all', force);
    }

    render() {
        // const {
        //     schemaId,
        //     version,
        //     compatibility,
        //     type: schemaType,
        //     schema,
        //     rawSchema
        // } = schemaDetails;

        // let { type, name, namespace, doc, fields } = schema as Schema;
        // if (schemaType === SchemaType.JSON) {
        //     const jsonSchema = schema as JsonSchema;
        //     ({ type, title: name, description: doc, $id: namespace } = jsonSchema);
        //     fields = [];
        //     if (jsonSchema.properties) {
        //         for (const p in jsonSchema.properties) {
        //             const property = jsonSchema.properties[p];
        //             fields.push(convertJsonField(p, property));
        //         }
        //     }
        // }

        // Create the object that will be shown in the JSON Viewer
        // From this new "display object" we can remove the 'rawSchema' that we added


        const schema = api.schemaDetails.get(this.subjectNameRaw);
        if (!schema) return DefaultSkeleton;

        // const versions = schemaDetails.registeredVersions ?? [];
        const defaultVersion = (typeof this.version == 'string')
            ? schema.versions[schema.versions.length - 1].version
            : this.version;

        return (
            <PageContent key="b">
                <Section py={4}>
                    <Flex>
                        <Statistic title="Type" value={schema.type}></Statistic>
                        <Statistic title="Subject" value={this.subjectNameRaw}></Statistic>
                        {/* <Statistic title="Schema ID" value={}></Statistic> */}
                        <Statistic title="Version" value={schema.latestActiveVersion}></Statistic>
                        <Statistic title="Compatibility" value={<span style={{ textTransform: 'capitalize' }}>{schema.compatibility.toLowerCase()}</span>}></Statistic>
                    </Flex>
                </Section>

                <Section>
                    <div style={{ display: 'flex', alignItems: 'flex-start', columnGap: '1.5em', marginBottom: '1em' }}>
                        <Label text="Version">
                            <Select style={{ minWidth: '200px' }}
                                defaultValue={defaultVersion}
                                onChange={(version) => {
                                    this.version = version as 'latest' | number;
                                    editQuery(x => {
                                        x.version = String(this.version);
                                    });
                                }}
                                disabled={schema.versions.length == 0}
                            >
                                {schema.versions.map(v => <Select.Option value={v.version} key={v.version}>Version {v.version}</Select.Option>)}
                            </Select>
                        </Label>

                        {/* <Label text="Details" style={{ alignSelf: 'stretch' }}>
                            <div style={{ display: 'inline-flex', flexWrap: 'wrap', minHeight: '32px', alignItems: 'center', rowGap: '.3em' }}>
                                {Object.entries({
                                    'Type': type,
                                    'Name': name,
                                    'Namespace': namespace,
                                }).map(([k, v]) => {
                                    if (!k || v === undefined || v === null) return null;
                                    return <Tag key={k}><span style={{ color: '#2d5b86' }}>{k}:</span> {toSafeString(v)}</Tag>
                                })}
                                {!!doc && <a href={doc}>
                                    <Tag style={{ cursor: 'pointer' }}><span style={{ color: '#2d5b86' }}>Documentation:</span> <a style={{ textDecoration: 'underline' }} href={doc}>{doc}</a></Tag>
                                </a>}
                            </div>

                        </Label> */}
                    </div>

                    <div style={{ marginBottom: '1.5em', display: 'flex', gap: '1em' }}>
                        <OptionGroup label=""
                            options={{
                                'Show Fields': 'fields',
                                'Show JSON': 'json',
                            }}
                            value={uiSettings.schemaDetails.viewMode}
                            onChange={s => uiSettings.schemaDetails.viewMode = s}
                        />

                        <NoClipboardPopover placement="top">
                            <div>
                                {' '}
                                {/* the additional div is necessary because popovers do not trigger on disabled elements, even on hover */}
                                <Tooltip label="Copy raw JSON to clipboard" placement="top" hasArrow={true}>
                                    <Button
                                        isDisabled={!isClipboardAvailable}
                                        onClick={() => {
                                            // navigator.clipboard.writeText(rawSchema);
                                            // message.success('Schema copied to clipboard', 1.2);
                                        }}
                                    >
                                        <Icon as={AiOutlineCopy} color="#555" width="18px" />
                                    </Button>
                                </Tooltip>
                            </div>
                        </NoClipboardPopover>

                    </div>

                    {/* <div>
                        {uiSettings.schemaDetails.viewMode == 'json' &&
                            <KowlJsonView
                                shouldCollapse={false}
                                collapsed={false}
                                src={jsonViewObject}
                                style={{
                                    border: 'solid thin lightgray',
                                    borderRadius: '.25em',
                                    padding: '1em 1em 1em 2em',
                                    marginBottom: '1.5rem',
                                }}
                            />
                        }

                        {uiSettings.schemaDetails.viewMode == 'fields' &&
                            <Table
                                size="small"
                                columns={[
                                    { title: 'Name', dataIndex: 'name', className: 'whiteSpaceDefault', }, // sorter: sortField('name')
                                    { title: 'Type', dataIndex: 'type', className: 'whiteSpaceDefault', render: renderSchemaType }, //  sorter: sortField('type'),
                                    { title: 'Default', dataIndex: 'default', className: 'whiteSpaceDefault' },
                                    { title: 'Documentation', dataIndex: 'doc', className: 'whiteSpaceDefault' },
                                ]}
                                rowKey="name"
                                dataSource={fields}
                                pagination={false}
                                style={{
                                    maxWidth: '100%',
                                    marginTop: '1.5rem',
                                    marginBottom: '1.5rem',
                                }}
                            />
                        }

                    </div> */}
                </Section>
            </PageContent>
        );
    }
}

function getVersionFromQuery(): 'latest' | number {
    const query = new URLSearchParams(window.location.search);
    if (query.has('version')) {
        const versionStr = query.get('version');

        if (versionStr != '' && !isNaN(Number(versionStr))) {
            return Number(versionStr);
        }

        if (versionStr == 'latest')
            return 'latest';

        console.log(`unknown version string in query: "${versionStr}" will be ignored, proceeding with "latest"`);
    }

    return 'latest';
}

export default SchemaDetailsView;

